package web

type WebEngine interface {
	addHandler(func())
}

type WebConfig struct {
	Addr string
}