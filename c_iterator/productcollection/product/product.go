package product

type Product struct {
	id   int
	name string
}

func NewProduct(id int, name string) *Product {
	return &Product{id: id, name: name}
}
