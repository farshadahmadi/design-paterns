package imagestorage

import (
	"fmt"

	"github.com/farshadahmadi/d_strategy/imagestorage/c_solution/compress"
	"github.com/farshadahmadi/d_strategy/imagestorage/c_solution/filter"
)

// This service is for storing images. As part of storage, this service should first compress the image with one of
// compression algorithms (PNG, JPEG, ...) and second filter the image with one of filtering algorithms (BlackAndWhite),
// and only then save the image (lets say to AWS S3 bucket). Solve the issue!

type ImageStorage struct {
	compressor compress.Compressor
	filter     filter.Filter
}

func NewImageStorage(compressor compress.Compressor, filter filter.Filter) *ImageStorage {
	return &ImageStorage{compressor: compressor, filter: filter}
}

func (is *ImageStorage) StoreImage(fileName string, data []byte) error {
	is.compressor.Compress(data)

	is.filter.Filter(data)

	// do the actual storage here
	fmt.Println("image is stored.")

	return nil
}