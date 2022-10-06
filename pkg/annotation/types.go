package annotation

import (
	"reflect"
)

type AnnotationParams map[string]any

type Annotation struct {
	PackageName    string
	AnnotationName string
	TargetType     reflect.Type
	TargetValue    reflect.Value
	Params         AnnotationParams
	RawData        string
}

type PackageAnnotation interface {
	PackageName() string
	PackageVariableAnnotations() []Annotation
	PackageFunctionAnnotations() []Annotation
	StructAnnotations() []Annotation
	StructFieldAnnotations() []Annotation
	StructMethodAnnotations() []Annotation
}

type AnnotationProcessor interface {
	AnnotationName() string
	ProcessAnnotation(annotation Annotation)
}
