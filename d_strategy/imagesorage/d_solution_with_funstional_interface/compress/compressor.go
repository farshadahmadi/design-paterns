package compress

type Compressor interface {
	Compress([]byte)
}

type CompressorFunc func(bytes []byte)

func (cf CompressorFunc) Compress(bytes []byte) {
	cf(bytes)
}
