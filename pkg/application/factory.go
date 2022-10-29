package application

import (
	"github.com/gowcar/gboot/pkg/ref"
	"reflect"
)

type GenericObjectFactory struct{}

func (f GenericObjectFactory) NewInstance(target *ref.TypeDesc) any {
	v :=  reflect.TypeOf(target.Object).Elem()
	return reflect.New(v).Interface()
}