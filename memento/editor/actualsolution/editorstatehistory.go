package actualsolution

type editorStateHistory struct {
	history []*editorState
}

func NewEditorStateHistory() *editorStateHistory {
	return &editorStateHistory{}
}

func (esh *editorStateHistory) PushState(state *editorState) {
	esh.history = append(esh.history, state)
}

func (esh *editorStateHistory) PopState() *editorState {
	s := esh.history[len(esh.history)-1]
	esh.history = esh.history[:len(esh.history)-1]
	return s
}
