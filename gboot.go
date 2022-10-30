package gboot

import (
	"context"
	"github.com/gowcar/gboot/pkg/annotation"
	"github.com/gowcar/gboot/pkg/application"
	"github.com/gowcar/gboot/pkg/config"
	"github.com/gowcar/gboot/pkg/log"
	"github.com/gowcar/gboot/pkg/web"
	"sync"
)

var once sync.Once

func ConfigGet(key string) any {
	return config.ConfigGet(key)
}

func StartApplication() {
	initialize(nil)
	log.Debug("application %v started", ConfigGet("application.name"))
}

func Waitfor() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	<- ctx.Done()
}

func RegisterAnnotations(packages []annotation.PackageAnnotation) {
	initialize(packages)
}

func initialize(packages []annotation.PackageAnnotation) {
	once.Do(func() {
		initConfig()
		initLogger()
		annotation.InitAnnotations(packages)
		initApplication()
		initWebEngine()
		startWebEngine()
	})
}

func startWebEngine() {
	web.Start()
}

func initWebEngine() {
	web.Initialize()
}

func initApplication() {
	log.Debug("GBoot application initializing")
	log.Debug("Web framework ====> %s", config.Config().Application.WebFramework)
	log.Debug("ORMapping framework ====> %s", config.Config().Application.DBFramework)
	application.Initialize()
}

func initConfig() {
	const fileName = "gboot.yml"
	config.InitConfig(fileName)
}

func initLogger() {
	options := log.DefaultOption()
	options.LogFolder = config.Config().Log.Folder
	options.LogFile = config.Config().Log.File
	options.LogLevel = config.Config().Log.Level
	log.InitLogger(options)
}
