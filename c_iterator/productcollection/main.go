package main

import (
	"fmt"

	"github.com/farshadahmadi/c_iterator/productcollection/product"
)

func main() {
	pc := product.NewCollection()
	pc.Add(product.NewProduct(1, "p1"))
	pc.Add(product.NewProduct(2, "p2"))
	pc.Add(product.NewProduct(3, "p3"))

	ci := pc.CreateIterator()

	for ci.Reset(); !ci.IsDone(); ci.Next() {
		fmt.Println(ci.CurrentItem())
	}
}
