package web

import (
	"github.com/gowcar/gboot/pkg/config"
	"sync"
)

var engine Engine
var once sync.Once

func Initialize()  {
	once.Do(func() {
		chooseEngine()
		engine.initialize()
	})
	AddHandler("GET", "/", func() {

	})
}

func Start()  {
	go engine.start()
}

func chooseEngine() {
	frm := config.Config().Application.WebFramework
	switch new(Framework).ValueOf(frm) {
	case Gin:
		engine = &GinEngine{}
	}
	engine = &FiberEngine{}
}

func AddHandler(method string, path string, handler any) {
	engine.registerHandler(method, path, handler)
}