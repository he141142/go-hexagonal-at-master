package constant

type EntityEnum int32

const (
	Form EntityEnum = iota + 1
	Todo
)

var EntitiesList = []string{
	"form",
	"todo",
}

func (et EntityEnum) String() string {
	return EntitiesList[et-1]
}

