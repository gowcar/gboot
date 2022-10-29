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
	PackageName       string
	AnnotationName    string
	AnnotationTarget  Target
	TargetTypeName    string
	TargetName        string
	TargetObject      any
	ParentAnnotations []Annotation
	Params            Params
	RawData           string
}

type PackageAnnotation interface {
	PackageName() string
	AllAnnotations() []Annotation
}

type Processor interface {
	Name() string
	AcceptTargets() Target
	Process(anno Annotation)
}
