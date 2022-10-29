package ref

import (
	"github.com/gowcar/gboot/pkg/log"
	"reflect"
	"unsafe"
)

var vars map[string]*TypeDesc
var funcs map[string]*TypeDesc
var structs map[string]*TypeDesc

func init() {
	vars = make(map[string]*TypeDesc)
	funcs = make(map[string]*TypeDesc)
	structs = make(map[string]*TypeDesc)
}

func GetFieldValue(obj any, fieldName string) any {
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}
	if fieldValue := objValue.FieldByName(fieldName); (fieldValue != reflect.Value{}) {
		fieldValue = reflect.NewAt(fieldValue.Type(), unsafe.Pointer(fieldValue.UnsafeAddr())).Elem()
		return GetValue(fieldValue)
	}
	return nil
}
func SetFieldValue(obj any, fieldName string, targetValue any) {
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}
	if fieldValue := objValue.FieldByName(fieldName); (fieldValue != reflect.Value{}) {
		fieldValue = reflect.NewAt(fieldValue.Type(), unsafe.Pointer(fieldValue.UnsafeAddr())).Elem()
		SetValue(fieldValue, targetValue)
	} else {
		log.Warn("field %v not found in %v", fieldName, objValue.Type().Name())
	}
}

func GetValue(targetObject reflect.Value) any {
	value := targetObject
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return value.Interface()
}

func SetValue(targetObject reflect.Value, targetValue any) {
	value := targetObject
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	t := reflect.ValueOf(targetValue)
	if value.CanSet() && t.CanConvert(value.Type()) {
		t = t.Convert(value.Type())
		value.Set(t)
	} else {
		log.Error("the value %v cannot be assign a value %v", value, targetValue)
	}
}

func RegisterVar(desc *TypeDesc) {
	vars[desc.Name] = desc
}

func RegisterFunc(desc *TypeDesc) {
	funcs[desc.Name] = desc
}

func RegisterStruct(desc *TypeDesc) {
	structs[desc.Name] = desc
}

func GetVar(name string) *TypeDesc {
	return vars[name]
}

func GetStruct(name string) *TypeDesc {
	return structs[name]
}

type TypeDesc struct {
	Name        string
	Object      any
	ProxyObject any
}
