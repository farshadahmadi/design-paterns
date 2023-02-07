package problem

// we want to implement undo mechanism for this editor

type Editor struct {
	content string
}

func NewEditor(content string) *Editor {
	return &Editor{content: content}
}

func (e *Editor) Content() string {
	return e.content
}

func (e *Editor) SetContent(content string) {
	e.content = content
}
