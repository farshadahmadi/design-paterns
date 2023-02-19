package imagestorage

import (
	"fmt"

	"github.com/farshadahmadi/d_strategy/imagestorage/b_badsolution/compress"
	"github.com/farshadahmadi/d_strategy/imagestorage/b_badsolution/filter"
)

// The trivial bad solution is to use enums. So client code will define the compressing and filtering algorithms (
// either via a member variables or by sending directly to StorageImage function). This solution has some problems:
// - Single Responsibility Principle: ImageStorage class should only do storage. Compression and filtering should go to other services
// - Open-Closed Principle: if later a new compression or filtering is introduced, then ImageStorage implementation must change!
//   However, ImageStorage should be open for extension, but closed for modification!

type ImageStorage struct {
	compress compress.Compress
	filter   filter.Filter
}

func NewImageStorage(compress compress.Compress, filter filter.Filter) *ImageStorage {
	return &ImageStorage{compress: compress, filter: filter}
}

func (is *ImageStorage) StoreImage(fileName string, data []byte) error {
	if is.compress == compress.PNG {
		fmt.Println("Image is compressed using,", compress.PNG)
	} else if is.compress == compress.JPEG {
		fmt.Println("Image is compressed using", compress.JPEG)
	}

	if is.filter == filter.BlackAndWhite {
		fmt.Println("Image is filtered using", filter.BlackAndWhite)
	}

	// do the actual storage here
	fmt.Println("image is stored.")

	return nil
}
