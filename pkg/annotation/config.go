package annotation

import (
	"github.com/gowcar/gboot/internal/ref"
	"github.com/gowcar/gboot/pkg/config"
	"github.com/gowcar/gboot/pkg/log"
)

type ConfigAnnotationProcessor struct {
}

func (ConfigAnnotationProcessor) AnnotationName() string {
	return "@Config"
}

func (ConfigAnnotationProcessor) ProcessAnnotation(anno Annotation) {
	ref.RegisterType(&ref.TypeDesc{
		Name:   anno.TargetName,
		Object: anno.TargetObject,
	})
	key := anno.Params["default"]
	configValue := config.ConfigGet(key.(string))
	if configValue != nil {
		ref.SetValue(anno.TargetObject, configValue)
		log.Debug("inject package variable %v", configValue)
	}
}
