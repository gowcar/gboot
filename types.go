package gboot

type (
	ORM          int8
)

const (
	GORM ORM = 1 << iota
	SDDL
)

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