package annotation

import (
	"github.com/gowcar/gboot/internal/ref"
	"github.com/gowcar/gboot/pkg/config"
	"github.com/gowcar/gboot/pkg/log"
)

type ConfigAnnotationProcessor struct {
}

func (p ConfigAnnotationProcessor) AcceptTargets() Target {
	return Var | StructField
}

func (p ConfigAnnotationProcessor) Process(anno Annotation) {
	key := anno.Params["default"]
	configValue := config.ConfigGet(key.(string))
	if configValue != nil {
		ref.SetValue(anno.TargetObject, configValue)
		log.Debug("inject package variable %v", configValue)
	}
}

func (p ConfigAnnotationProcessor) Name() string {
	return "@Config"
}