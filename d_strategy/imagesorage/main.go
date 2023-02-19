package main

import (
	"fmt"

	"github.com/farshadahmadi/d_strategy/imagestorage/b_badsolution/compress"
	"github.com/farshadahmadi/d_strategy/imagestorage/b_badsolution/filter"
	"github.com/farshadahmadi/d_strategy/imagestorage/b_badsolution/imagestorage"
	bettercompress "github.com/farshadahmadi/d_strategy/imagestorage/c_solution/compress"
	betterfilter "github.com/farshadahmadi/d_strategy/imagestorage/c_solution/filter"

	betterimagestorage "github.com/farshadahmadi/d_strategy/imagestorage/c_solution/imagestorage"
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

	fmt.Println("--------")

	// The good solution is to use polymorphism and composition/delegation. This solves the two aforementioned issues.
	// A compressor interface is defined. Every concrete compressor will implement that interface. The same goes for filtering.
	// ImageStorage class will have member variables of interface types. Based on the given concrete implementations, Image Storage class
	// can change the compression or filtering algorithms at run time.
	// Other approach is to remove member variables, and give concrete strategies directly to the image storage function:
	// is1.StoreImage("image1", bettercompress.NewJpegCompressor(), betterfilter.NewBlackAndWhiteFilter())

	is1 := betterimagestorage.NewImageStorage(bettercompress.NewJpegCompressor(), betterfilter.NewBlackAndWhiteFilter())
	err = is1.StoreImage("image1", []byte{})
	if err != nil {
		fmt.Println("cannot store image", err)

		return
	}
}
