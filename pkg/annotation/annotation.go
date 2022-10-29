package annotation

import (
	"github.com/gowcar/gboot/internal/ref"
	"github.com/gowcar/gboot/pkg/log"
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
	log.Debug("pkg.name is %v", pkg.PackageName())
	for _, anno := range pkg.AllAnnotations() {
		registerTypes(anno)
		p, _ := processors[anno.AnnotationName]
		if p != nil  && anno.AnnotationTarget & p.AcceptTargets() == anno.AnnotationTarget {
			p.Process(anno)
		}
	}
}

func registerTypes(anno Annotation) {
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
	}

}

func processPackageVariableAnnotation(anno Annotation) {
	p, _ := processors[anno.AnnotationName]
	if p != nil  && anno.AnnotationTarget & p.AcceptTargets() == anno.AnnotationTarget {
		p.Process(anno)
	}
}

func processPackageFunctionAnnotation(anno Annotation) {
	p, _ := processors[anno.AnnotationName]
	if p != nil {
		p.Process(anno)
	}
}
