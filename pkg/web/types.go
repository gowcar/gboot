package web

type (
	Framework int8
)
const (
	Fiber Framework = 1 << iota
	Gin
)
func (p Framework) ValueOf(v string) Framework{
	switch v {
	case "Fiber":
		return Fiber
	case "Gin":
		return Gin
	}
	return 0
}
func (p Framework) String() string {
	switch p {
	case Fiber:
		return "Fiber"
	case Gin:
		return "Gin"
	default:
		return "UNKNOWN"
	}
}


type Engine interface {
	initial()
	registerHandler()
	start()
}

type EngineConfig struct {
	Addr string
}