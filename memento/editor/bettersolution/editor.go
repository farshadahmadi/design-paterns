package bettersolution

type Editor struct {
	history *editorStateHistory
}

func NewEditor() *Editor {
	return &Editor{
		history: NewEditorStateHistory(),
	}
}

func (e *Editor) SetContent(content string) {
	state := &editorState{
		content: content,
	}
	e.history.pushState(state)
}

func (e *Editor) Undo() {
	e.history.popState()
}

func (e *Editor) Content() string {
	state := e.history.getState()
	return state.content
}
