package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gowcar/gboot/pkg/config"
)

type FiberEngine struct{
	app *fiber.App
}

func (engine *FiberEngine) initial() {
	engine.app = fiber.New()
}

func (engine *FiberEngine) registerHandler() {
	engine.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}

func (engine *FiberEngine) start() {
	engine.app.Listen(config.Config().Application.Addr)
}
