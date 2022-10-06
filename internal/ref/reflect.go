package ref

import (
	"fmt"
	"github.com/gowcar/gboot/pkg/log"
	"reflect"
	"unsafe"
)

func GetValue(value reflect.Value) any {
	return value.Interface()
}

func SetValue(value reflect.Value, targetValue any, targetType reflect.Type) {
	t := reflect.ValueOf(targetValue)
	if value.CanSet() && t.CanConvert(targetType) {
		t = t.Convert(targetType)
		value.Set(t)
	} else {
		log.Error("the value %v cannot be assign a value %v", value, targetValue)
	}
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
