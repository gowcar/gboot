package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gowcar/gboot/pkg/config"
	"github.com/gowcar/gboot/pkg/log"
)

type FiberEngine struct{
	app *fiber.App
}

func (engine *FiberEngine) initialize() {
	engine.app = fiber.New()
}

func (engine *FiberEngine) registerHandler(method string, path string, fn any) {
	engine.app.Add(method, path, proxyFunc(fn))
}

func (engine *FiberEngine) start() {
	//var m = make(map[string]fiber.Handler)
	//
	//engine.app.Get("*", func(c *fiber.Ctx) error {
	//	h, exist := m[c.Path()]
	//	if !exist {
	//		c.SendStatus(fiber.StatusNotFound)
	//	} else {
	//		return h(c)
	//	}
	//	return nil
	//})
	engine.app.Listen(config.Config().Application.Addr)
	log.Debug("handlers : %v", engine.app.Stack())
}

func proxyFunc(fn any) fiber.Handler {
	return func(c *fiber.Ctx) error {
		before(c)
		err := c.SendString("Hello, World!")
		after(c)
		return err
	}
}



func after(c *fiber.Ctx) {
	log.Debug("after happend")
}

func before(c *fiber.Ctx) {
	log.Debug("before happended")
}