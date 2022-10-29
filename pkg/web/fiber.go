package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gowcar/gboot/pkg/log"
	"sync"
)

var webApp *fiber.App
var config *WebConfig
var once sync.Once
var handlers = make(map[string]fiber.Handler)

func Initialize(webConf *WebConfig) *fiber.App{
	once.Do(func() {
		config = webConf
		webApp = fiber.New()
	})
	return webApp
}

func AddHandler(path string, handler fiber.Handler) {
	//webApp.Add("GET", path, handler)
	handlers[path] = handler
	//webApp.Get(path, handler)
}

func Start()  {
	webApp.All("*", func(ctx *fiber.Ctx) error {
		log.Debug("path ==> %v", ctx.Path())
		h, exist := handlers[ctx.Path()]
		if !exist {
			ctx.SendStatus(fiber.StatusNotFound)
		} else {
			h(ctx)
		}
		return nil
	})
	go webApp.Listen(config.Addr)
}
