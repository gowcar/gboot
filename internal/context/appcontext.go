package context

import "sync"

type ApplicationContext struct {
}

var instance *ApplicationContext
var once sync.Once

func Instance() *ApplicationContext {
	once.Do(func() {
		instance = &ApplicationContext{}
	})
	return instance
}

func (appContext *ApplicationContext) GetComponent(key string) any {
	return nil
}

func (appContext *ApplicationContext) SayHello(key string) string {
	return "Hello " + key
}
