package proxy

import (
	"github.com/gowcar/gboot/pkg/log"
	"reflect"
)

type InvocationMethod struct {
	Name string
	Type reflect.Type
}

func (im InvocationMethod) Invoke(obj any, args []reflect.Value) []reflect.Value {
	v := reflect.ValueOf(obj).MethodByName(im.Name)

	if !v.IsValid() {
		panic("Can not found method " + im.Name + " in " + reflect.ValueOf(obj).Type().String())
	}

	return v.Call(args)
}

var InvocationProxy = invocationProxy{}

type invocationProxy struct {
}

type InvocationHandler func(obj any, method InvocationMethod, args []reflect.Value) []reflect.Value

func (ip invocationProxy) NewProxyInstance(itf any, handler InvocationHandler) any {

	t := reflect.TypeOf(itf)

	if t.Kind() != reflect.Ptr {
		panic("Need a pointer of interface struct")
	}

	if t.Elem().Kind() != reflect.Struct {
		panic("Need a pointer of interface struct")
	}

	t = t.Elem()
	ot := reflect.ValueOf(itf).Elem()
	n := ot.NumField()

	for idx := 0; idx < n; idx++ {
		f := t.Field(idx)
		of := ot.Field(idx)

		if f.Type.Kind() == reflect.Func {

			if !of.CanSet() {
				log.Debug("field %v is a readonly func, ignore proxy", f.Name)
				continue
			}

			log.Debug("field %v is a func, proxy now", f.Name)

			target := reflect.MakeFunc(of.Type(), func(args []reflect.Value) (results []reflect.Value) {
				//param := make([]any, 0, len(args))
				//
				//for _, o := range args {
				//	param = append(param, o.Interface())
				//}

				method := InvocationMethod{Name: f.Name, Type: f.Type}

				rtn := handler(itf, method, args)

				//if rtn == nil {
				//	return []reflect.Value{}
				//}
				//
				//rtnV := make([]reflect.Value, 0, len(rtn))
				//
				//for _, o := range rtn {
				//	rtnV = append(rtnV, reflect.ValueOf(o))
				//}

				return rtn
			})
			of.Set(target)
		} else {
			log.Debug("field %v is not a func, ignore proxy", f.Name)
		}
	}

	return itf
}
