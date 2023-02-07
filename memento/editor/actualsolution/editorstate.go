package actualsolution

type editorState struct {
	content string
}

func newEditorState(content string) *editorState {
	return &editorState{content: content}
}
