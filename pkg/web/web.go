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
	})
}

func chooseEngine() {
	frm := config.Config().Application.WebFramework
	switch new(Framework).ValueOf(frm) {
	case Gin:
		engine = &GinEngine{}
	}
	engine = &FiberEngine{}
}

func Start()  {
	engine.start()
}