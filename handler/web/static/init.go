package static

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

type (
	HtmlTemplate struct {
		app *fiber.App
	}
)

func New(app *fiber.App) *HtmlTemplate {
	return &HtmlTemplate{
		app: app,
	}
}

func (h *HtmlTemplate) Register() {
	h.app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("file_upload", fiber.Map{
			"Token": os.Getenv("AUTH_TOKEN"),
		})
		//return c.SendFile("./handler/web/static/file _upload.html",)
	})
}
