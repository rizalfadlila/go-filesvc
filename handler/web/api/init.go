package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizalfadlila/go-filesvc/model"
	"github.com/rizalfadlila/go-filesvc/usecase"
	"io"
	"os"
)

type (
	Api struct {
		app          *fiber.App
		imageUsecase usecase.ImageUsecase
	}

	Opts struct {
		App          *fiber.App
		ImageUsecase usecase.ImageUsecase
	}
)

func New(opts Opts) *Api {
	return &Api{
		app:          opts.App,
		imageUsecase: opts.ImageUsecase,
	}
}

func (a *Api) Register() {
	a.app.Post("/upload", func(c *fiber.Ctx) error {
		if c.FormValue("token") != os.Getenv("AUTH_TOKEN") {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "invalid creds",
			})
		}

		// Parse the form file
		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "File upload failed",
				"error":   err.Error(),
			})
		}

		// Reject files larger than 8 megabytes
		if file.Size > 8<<20 { // 8 megabytes in bytes
			return c.Status(fiber.StatusRequestEntityTooLarge).JSON(fiber.Map{
				"message": "File size exceeds the limit of 8 megabytes",
			})
		}

		// Open the uploaded file
		src, err := file.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to open uploaded file",
				"error":   err.Error(),
			})
		}
		defer src.Close()

		// Create a temporary file
		tempFile, err := os.CreateTemp("", "uploaded-*.tmp")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to create temporary file",
				"error":   err.Error(),
			})
		}
		defer tempFile.Close()

		// Copy the uploaded file to the temporary file
		_, err = io.Copy(tempFile, src)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to copy file to temporary file",
				"error":   err.Error(),
			})
		}

		image := &model.ImageMetadata{
			Filename: file.Filename,
			Filepath: tempFile.Name(),
			Size:     file.Size,
		}
		if err := a.imageUsecase.Save(c.UserContext(), image); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to save file metadata",
				"error":   err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message":  "File uploaded successfully",
			"filepath": tempFile.Name(),
		})
	})
}
