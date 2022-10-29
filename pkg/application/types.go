package application

import (
	"github.com/gowcar/gboot/pkg/ref"
)


type component struct {
	componentType *ref.TypeDesc
	instance      any
}

type Injector interface {
	Inject(target *ref.TypeDesc)
}

type InstanceFactory interface {
	NewInstance(target *ref.TypeDesc) any
}