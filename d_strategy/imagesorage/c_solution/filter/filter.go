package filter

type Filter interface {
	Filter([]byte)
}
