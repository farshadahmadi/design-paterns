package a_problem

type Tool int

const (
	Undefined Tool = iota
	Selection
	Brush
	Eraser
)

func Tools() []Tool {
	return []Tool{Selection, Brush, Eraser}
}

func (t Tool) IsValid() bool {
	for _, tool := range Tools() {
		if t == tool {
			return true
		}
	}

	return false
}

func (t Tool) String() string {
	switch t {
	case Undefined:
		return "Undefined"
	case Selection:
		return "Selection"
	case Brush:
		return "Brush"
	case Eraser:
		return "Eraser"
	}

	return "Unknown"
}
