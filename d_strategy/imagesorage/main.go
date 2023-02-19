package main

import (
	"fmt"

	"github.com/farshadahmadi/d_strategy/imagestorage/a_problem/compress"
	"github.com/farshadahmadi/d_strategy/imagestorage/a_problem/filter"
	"github.com/farshadahmadi/d_strategy/imagestorage/b_badsolution/imagestorage"
)

func main() {
	// The trivial bad solution is to use enums. So client code will define the compressing and filtering algorithms (
	// either via a member variables or by sending directly to StorageImage function). This solution has some problems:
	// - Single Responsibility Principle: ImageStorage class should only do storage. Compression and filtering should go to other services
	// - Open-Closed Principle: if later a new compression or filtering is introduced, then ImageStorage implementation must change!
	//   However, ImageStorage should be open for extension, but closed for modification!

	is := imagestorage.NewImageStorage(compress.JPEG, filter.BlackAndWhite)
	err := is.StoreImage("image1", []byte{})
	if err != nil {
		fmt.Println("cannot store image", err)

		return
	}
}
