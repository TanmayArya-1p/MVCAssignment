package types

const (
	DefaultLimit  = -1
	DefaultOffset = 0
)

type Page struct {
	Limit  int
	Offset int
}
