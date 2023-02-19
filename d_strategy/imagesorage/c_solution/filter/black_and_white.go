package filter

import (
	"fmt"
)

type BlackAndWhiteFilter struct{}

func NewBlackAndWhiteFilter() *BlackAndWhiteFilter {
	return &BlackAndWhiteFilter{}
}

func (bwf *BlackAndWhiteFilter) Filter(bytes []byte) {
	fmt.Println("Image is filtered using Black And White.")
}
