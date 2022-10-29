package annotation

import (
	"github.com/gowcar/gboot/pkg/application"
	"github.com/gowcar/gboot/pkg/log"
	"github.com/gowcar/gboot/pkg/ref"
)

var processors map[string]Processor

func InitAnnotations(packages []PackageAnnotation) {
	internalProcessors()
	for _, pkg := range packages {
		processPackage(pkg)
	}
}

func internalProcessors() {
	processors = make(map[string]Processor)
	processorsList := []Processor{
		&ConfigAnnotationProcessor{},
	}
	for _, processor := range processorsList {
		RegisterProcessor(processor)
	}
}

func RegisterProcessor(processor Processor) {
	p, _ := processors[processor.Name()]
	if p == nil {
		processors[processor.Name()] = processor
	}
}

func processPackage(pkg PackageAnnotation) {
	for _, anno := range pkg.AllAnnotations() {
		prepare(anno, pkg.AllAnnotations())
		p, exist := processors[anno.AnnotationName]
		if exist && p != nil  && anno.AnnotationTarget & p.AcceptTargets() == anno.AnnotationTarget {
			p.Process(anno)
		} else {
			log.Warn("No processor found for the annotation:%v", anno.AnnotationName)
		}
	}
}

func prepare(anno Annotation, all []Annotation) {
	switch anno.AnnotationTarget {
	case Var:
		ref.RegisterVar(&ref.TypeDesc{
			Name:        anno.TargetName,
			Object:      anno.TargetObject,
		})
	case Func:
		ref.RegisterFunc(&ref.TypeDesc{
			Name:        anno.TargetName,
			Object:      anno.TargetObject,
		})
	case Struct:
		ref.RegisterStruct(&ref.TypeDesc{
			Name:        anno.TargetName,
			Object:      anno.TargetObject,
		})
		application.RegisterInstanceFactory(anno.TargetName, &application.GenericObjectFactory{})
	}
	anno.ParentAnnotations = findParentAnnotation(anno, all)
}

func findParentAnnotation(anno Annotation, all []Annotation) []Annotation {
	parent := make([]Annotation, 0)
	for _, item := range all {
		if anno.TargetTypeName == item.TargetName {
			parent = append(parent, item)
		}
	}
	return parent
}