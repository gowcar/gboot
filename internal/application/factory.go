package application

import (
	"github.com/gowcar/gboot/internal/ref"
	"reflect"
)

type GenericObjectFactory struct{}

func (GenericObjectFactory) newInstance(target ref.TypeDesc) any {
	v :=  reflect.TypeOf(target.Object)
	return reflect.New(v).Interface()
}