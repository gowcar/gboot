package application

import (
	"github.com/gowcar/gboot/pkg/ref"
)

var varInjectors map[string]Injector
var fieldInjectors map[string]map[string]Injector
var instanceCreators map[string]InstanceFactory

var container map[string]*component

func init() {
	varInjectors = make(map[string]Injector)
	fieldInjectors = make(map[string]map[string]Injector)
	instanceCreators = make(map[string]InstanceFactory)
	container = make(map[string]*component)
}

func Initialize() {
	injectVars()
	createInstances()
	injectStructFields()
}

func injectStructFields() {
	for k, comp:= range container{
		injectors, exist:= fieldInjectors[k]
		if exist {
			for fieldName, injector:= range injectors {
				injector.Inject(&ref.TypeDesc{
					Name:        fieldName,
					Object:      comp.instance,
				})
			}
		}
	}

}

func createInstances() {
	for s, fac := range instanceCreators {
		t := ref.GetStruct(s)
		obj := fac.NewInstance(t)
		container[s] = &component{
			componentType: t,
			instance:      obj,
		}
	}
}

func injectVars() {
	for k, injector := range varInjectors {
		v := ref.GetVar(k)
		injector.Inject(v)
	}
}

func GetComponent(componentId string) any {
	c, exist := container[componentId]
	if exist {
		return c.instance
	} else {
		return nil
	}
}

func RegisterVarInjector(varName string, injector Injector) {
	varInjectors[varName] = injector
}

func RegisterStructFieldInjector(structName string, fieldName string, injector Injector) {
	injectors, exist := fieldInjectors[structName]
	if !exist {
		injectors = make(map[string]Injector)
	}
	injectors[fieldName] = injector
	fieldInjectors[structName] = injectors
}

func RegisterInstanceFactory(structName string, factory InstanceFactory) {
	instanceCreators[structName] = factory
}

