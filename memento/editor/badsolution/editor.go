package badsolution

type Editor struct {
	history      []string
	historyIndex int
}

func NewEditor(content string) *Editor {
	return &Editor{
		history:      []string{content},
		historyIndex: 0,
	}
}

func (e *Editor) Content() string {
	return e.history[e.historyIndex]
}

func (e *Editor) SetContent(content string) {
	e.historyIndex++
	e.history = append(e.history, content)
}

func (e *Editor) Undo() {
	if e.historyIndex > 0 {
		e.historyIndex--
	}
}
