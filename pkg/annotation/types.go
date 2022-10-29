package annotation

type Target int8

const (
	Var = 1 << iota
	Func
	Struct
	StructField
	StructMethod
	Interface
	InterfaceMethod
)

const AllTarget Target = Var|Func|Struct|StructField|StructMethod|Interface|InterfaceMethod

type Params map[string]any

type Annotation struct {
	PackageName      string
	AnnotationName   string
	AnnotationTarget Target
	TargetName       string
	TargetObject any
	Params       Params
	RawData      string
}

type PackageAnnotation interface {
	PackageName() string
	AllAnnotations() []Annotation
	//PackageVariableAnnotations() []Annotation
	//PackageFunctionAnnotations() []Annotation
	//StructAnnotations() []Annotation
	//StructFieldAnnotations() []Annotation
	//StructMethodAnnotations() []Annotation
}

type Processor interface {
	Name() string
	AcceptTargets() Target
	Process(anno Annotation)
}
