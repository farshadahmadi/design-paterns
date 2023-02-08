package editorandstate

type EditorState interface {
	Content() string
}

type State struct {
	content string
}

func (e *State) Content() string {
	return e.content
}

func newEditorState(content string) *State {
	return &State{content: content}
}
