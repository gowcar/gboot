package application

var varInjectors map[string]Injector
var fieldInjectors map[string]Injector
var instanceCreators map[string]InstanceFactory

var container map[string]component

func Initialize() {

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
	fieldInjectors[structName+"."+fieldName] = injector
}

func RegisterInstanceFactory(structName string, factory InstanceFactory) {
	instanceCreators[structName] = factory
}

