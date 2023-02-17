package product

type CollectionIterator struct {
	collection *Collection
	index      int
}

func NewCollectionIterator(collection *Collection) *CollectionIterator {
	return &CollectionIterator{
		collection: collection,
		index:      0,
	}
}

func (c *CollectionIterator) IsDone() bool {
	return c.index >= len(c.collection.products)
}

func (c *CollectionIterator) CurrentItem() *Product {
	return c.collection.products[c.index]
}

func (c *CollectionIterator) Next() {
	c.index++
}

func (c *CollectionIterator) Reset() {
	c.index = 0
}
