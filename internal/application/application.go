package application

import "github.com/gowcar/gboot/internal/ref"

var container map[string]component

type component struct {
	componentType ref.TypeDesc
	instance      any
}

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
