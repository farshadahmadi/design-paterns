package filter

type Filter int

const (
	Undefined Filter = iota
	BlackAndWhite
)

func (c Filter) String() string {
	switch c {
	case Undefined:
		return "Undefined"
	case BlackAndWhite:
		return "Black And White"
	}

	return "Unknown"
}