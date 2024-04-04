package web

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/rizalfadlila/go-filesvc/config"
	"github.com/rizalfadlila/go-filesvc/handler/web/api"
	"github.com/rizalfadlila/go-filesvc/handler/web/static"
	"github.com/rizalfadlila/go-filesvc/usecase"
)

type (
	module struct {
		config       config.Main
		app          *fiber.App
		imageUsecase usecase.ImageUsecase
	}

	WebHanlder interface {
		Run() error
	}

	WebOpts struct {
		ImageUsecase usecase.ImageUsecase
		Config       config.Main
	}
)

func NewWebHandler(opts WebOpts) WebHanlder {
	return &module{
		config:       opts.Config,
		imageUsecase: opts.ImageUsecase,
		app: fiber.New(
			fiber.Config{
				Views: html.New("./handler/web/static/", ".html"),
			},
		),
	}
}

func (m *module) Run() error {
	m.app.Use(logger.New())
	m.app.Use(recover.New())

	static.New(m.app).Register()
	api.New(api.Opts{
		App:          m.app,
		ImageUsecase: m.imageUsecase,
	}).Register()

	return m.app.Listen(fmt.Sprintf(":%v", m.config.WebServer.Port))
}
