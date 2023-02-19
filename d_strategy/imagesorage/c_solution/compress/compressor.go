package compress

type Compressor interface {
	Compress([]byte)
}
