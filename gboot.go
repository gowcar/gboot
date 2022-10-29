package gboot

import (
	"github.com/gowcar/gboot/pkg/annotation"
	"github.com/gowcar/gboot/pkg/application"
	"github.com/gowcar/gboot/pkg/config"
	"github.com/gowcar/gboot/pkg/log"
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

func RegisterAnnotations(packages []annotation.PackageAnnotation) {
	initialize(packages)
}

func initialize(packages []annotation.PackageAnnotation) {
	once.Do(func() {
		initConfig()
		initLogger()
		annotation.InitAnnotations(packages)
		initApplication()
	})
}

func initApplication() {
	log.Debug("GBoot application initializing")
	log.Debug("Web framework ====> %s", Fiber)
	log.Debug("ORMapping framework ====> %s", GORM)
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
