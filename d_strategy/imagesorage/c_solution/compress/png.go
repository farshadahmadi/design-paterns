package compress

import (
	"fmt"
)

type PngCompressor struct{}

func NewPngCompressor() *PngCompressor {
	return &PngCompressor{}
}

func (p *PngCompressor) Compress(bytes []byte) {
	fmt.Println("Image is compressed using PNG.")
}
