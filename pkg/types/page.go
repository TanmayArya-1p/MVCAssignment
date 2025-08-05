package types

const (
	DefaultLimit  = 1000000000000000000
	DefaultOffset = 0
)

type Page struct {
	Limit  int
	Offset int
}
