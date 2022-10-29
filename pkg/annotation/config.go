package annotation

import (
	"github.com/gowcar/gboot/pkg/application"
	"github.com/gowcar/gboot/pkg/config"
	"github.com/gowcar/gboot/pkg/log"
	"github.com/gowcar/gboot/pkg/ref"
	"reflect"
)

type ConfigAnnotationProcessor struct {
	anno Annotation
}

func (p ConfigAnnotationProcessor) Inject(target *ref.TypeDesc) {
	key := p.anno.Params["default"]
	configValue := config.ConfigGet(key.(string))
	if configValue != nil {
		switch p.anno.AnnotationTarget {
		case Var:
			ref.SetValue(reflect.ValueOf(target.Object), configValue)
		case StructField:
			ref.SetFieldValue(target.Object, target.Name, configValue)
		}
		log.Debug("inject package variable %v", configValue)
	}
}

func (p ConfigAnnotationProcessor) AcceptTargets() Target {
	return Var | StructField
}

func (p ConfigAnnotationProcessor) Process(anno Annotation) {
	p.anno = anno
	switch anno.AnnotationTarget {
	case Var:
		application.RegisterVarInjector(anno.TargetName, p)
	case StructField:
		application.RegisterStructFieldInjector(anno.TargetTypeName, anno.TargetName, p)
	}
}

func (p ConfigAnnotationProcessor) Name() string {
	return "@Config"
}