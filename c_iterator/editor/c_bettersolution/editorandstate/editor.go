package editorandstate

type Editor struct {
	content string
}

func NewEditor() *Editor {
	return &Editor{}
}

func (e *Editor) Content() string {
	return e.content
}

func (e *Editor) SetContent(content string) {
	e.content = content
}

func (e *Editor) GetState() EditorState {
	return newEditorState(e.content)
}

func (e *Editor) SetState(state EditorState) {
	s := state.(*State)
	e.content = s.content
}
