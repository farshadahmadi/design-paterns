package compress

import (
	"fmt"
)

type JpegCompressor struct{}

func NewJpegCompressor() *JpegCompressor {
	return &JpegCompressor{}
}

func (p *JpegCompressor) Compress(bytes []byte) {
	fmt.Println("Image is compressed using JPEG.")
}
