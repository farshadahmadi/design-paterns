package editorsolution

type editorStateHistory struct {
	history []*editorState
	index   int
}

func NewEditorStateHistory() *editorStateHistory {
	return &editorStateHistory{
		index: -1,
	}
}

func (esh *editorStateHistory) pushState(state *editorState) {
	esh.index++
	esh.history = append(esh.history, state)
}

func (esh *editorStateHistory) popState() *editorState {
	esh.index--
	return esh.history[esh.index]
}

func (esh *editorStateHistory) getState() *editorState {
	return esh.history[esh.index]
}
