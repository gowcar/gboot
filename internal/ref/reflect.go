package ref

import (
	"fmt"
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

func GetValue(targetObject any) any {
	value := reflect.ValueOf(targetObject).Elem()
	return value.Interface()
}

func SetValue(targetObject any, targetValue any) {
	t := reflect.ValueOf(targetValue)
	value := reflect.ValueOf(targetObject).Elem()
	targetType := reflect.TypeOf(targetObject).Elem()
	if value.CanSet() && t.CanConvert(targetType) {
		t = t.Convert(targetType)
		value.Set(t)
	} else {
		log.Error("the value %v cannot be assign a value %v", value, targetValue)
	}
}

func NewInstance(name string) any {
	return nil
}

func NewProxyInstance(name string) any {
	return nil
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

type TypeDesc struct {
	Name        string
	Object      any
	ProxyObject any
}

func GetUnExportFiled(s interface{}, field string) (accessableField, addressableSource reflect.Value) {
	v := reflect.ValueOf(s)

	// 创建一个指向新地址的 addressableSource, 由于这个原因, 无法修改 s 本身的字段的值
	addressableSource = reflect.New(v.Type()).Elem()
	//addressableSource.Set(reflect.ValueOf(s))

	// 使用指针的方式一样
	accessableField = addressableSource.FieldByName(field)
	accessableField = reflect.NewAt(accessableField.Type(), unsafe.Pointer(accessableField.UnsafeAddr())).Elem()

	return accessableField, addressableSource
}

func SetUnExportFiled(s interface{}, filed string, val interface{}) error {
	v, addressableSource := GetUnExportFiled(s, filed)
	rv := reflect.ValueOf(val)
	if v.Kind() != v.Kind() {
		return fmt.Errorf("invalid kind, expected kind: %v, got kind:%v", v.Kind(), rv.Kind())
	}

	addressableSource.Set(rv)
	return nil
}
