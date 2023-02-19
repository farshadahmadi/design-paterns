package compress

type Compress int

const (
	Undefined Compress = iota
	PNG
	JPEG
)

func (c Compress) String() string {
	switch c {
	case Undefined:
		return "Undefined"
	case PNG:
		return "PNG"
	case JPEG:
		return "JPEG"
	}

	return "Unknown"
}