package application

import "github.com/gowcar/gboot/internal/ref"


type component struct {
	componentType ref.TypeDesc
	instance      any
}

type Injector interface {
	Inject(target ref.TypeDesc)
}

type InstanceFactory interface {
	newInstance(target ref.TypeDesc) any
}