package imagestorage

// This service is for storing images. As part of storage, this service should first compress the image with one of
// compression algorithms (PNG, JPEG, ...) and second filter the image with one of filtering algorithms (BlackAndWhite),
// and only then save the image (lets say to AWS S3 bucket). Solve the issue!

type ImageStorage struct{}

func (is *ImageStorage) StoreImage(fileName string, data []byte) error {
	return nil
}
