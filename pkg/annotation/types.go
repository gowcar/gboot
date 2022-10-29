package annotation

type AnnotationParams map[string]any

type Annotation struct {
	PackageName    string
	AnnotationName string
	TargetName     string
	TargetObject   any
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
