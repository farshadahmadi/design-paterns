package product

import "github.com/farshadahmadi/c_iterator/productcollection/iterator"

type Collection struct {
	products []*Product
}

func NewCollection() *Collection {
	return &Collection{products: make([]*Product, 0)}
}

func (pc *Collection) CreateIterator() iterator.Iterator[*Product] {
	return NewCollectionIterator(pc)
}

func (pc *Collection) Add(p *Product) {
	pc.products = append(pc.products, p)
}
