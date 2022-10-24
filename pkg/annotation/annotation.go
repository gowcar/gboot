package annotation

import "github.com/gowcar/gboot/pkg/log"

var processors map[string]AnnotationProcessor

func InitAnnotations(packages []PackageAnnotation) {
	internalProcessors()
	for _, pkg := range packages {
		processPackage(pkg)
	}
}

func internalProcessors() {
	processors = make(map[string]AnnotationProcessor)
	processorsList := []AnnotationProcessor {
		&ConfigAnnotationProcessor{},
	}
	for _, processor := range processorsList {
		RegisterProcessor(processor.AnnotationName(), processor)
	}
}

func RegisterProcessor(name string, processor AnnotationProcessor) {
	p, _ := processors[name]
	if p == nil {
		processors[name] = processor
	}
}

func processPackage(pkg PackageAnnotation) {
	log.Debug("pkg.name is %v", pkg.PackageName())
	for _, anno := range pkg.PackageVariableAnnotations() {
		processPackageVariableAnnotation(anno)
	}
}

func processPackageVariableAnnotation(anno Annotation) {
	p, _:= processors[anno.AnnotationName]
	if p != nil {
		p.ProcessAnnotation(anno)
	}

}
