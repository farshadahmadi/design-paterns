package history

import "github.com/farshadahmadi/memento/editor/e_bestsolution/editorandstate"

type editorStateHistory struct {
	history []editorandstate.EditorState
}

func NewEditorStateHistory() *editorStateHistory {
	return &editorStateHistory{}
}

func (esh *editorStateHistory) PushState(state editorandstate.EditorState) {
	esh.history = append(esh.history, state)
}

func (esh *editorStateHistory) PopState() editorandstate.EditorState {
	s := esh.history[len(esh.history)-1]
	esh.history = esh.history[:len(esh.history)-1]
	return s
}
