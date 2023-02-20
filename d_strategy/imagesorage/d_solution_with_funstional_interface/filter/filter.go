package filter

type Filter interface {
	Filter(file []byte)
}

type FilterFunc func([]byte)

func (ff FilterFunc) Filter(file []byte) {
	ff(file)
}
