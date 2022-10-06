package gboot

import (
	"github.com/gowcar/gboot/internal/web"
)

type (
	WebFramework int8
	ORM          int8
)

const (
	Fiber WebFramework = 1 << iota
	GIN
)

const (
	GORM ORM = 1 << iota
	SDDL
)

func (p WebFramework) String() string {
	switch p {
	case Fiber:
		return "Fiber"
	case GIN:
		return "GIN"
	default:
		return "UNKNOWN"
	}
}
func (p ORM) String() string {
	switch p {
	case GORM:
		return "GORM"
	case SDDL:
		return "SDDL"
	default:
		return "UNKNOWN"
	}
}

type Context web.Context
