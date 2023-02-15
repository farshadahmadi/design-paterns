package history

import (
	"github.com/farshadahmadi/c_iterator/editor/a_problem/editorandstate"
	"github.com/farshadahmadi/c_iterator/editor/c_bettersolution/iterator"
)

type EditorStateHistory struct {
	history []editorandstate.EditorState
}

func NewEditorStateHistory() *EditorStateHistory {
	esh := &EditorStateHistory{
		history: make([]editorandstate.EditorState, 0),
	}

	return esh
}

func (esh *EditorStateHistory) CreateIterator() iterator.Iterator[editorandstate.EditorState] {
	return NewEditorHistoryIterator(esh)
}

func (esh *EditorStateHistory) PushState(state editorandstate.EditorState) {
	esh.history = append(esh.history, state)
}

func (esh *EditorStateHistory) PopState() editorandstate.EditorState {
	s := esh.history[len(esh.history)-1]
	esh.history = esh.history[:len(esh.history)-1]
	return s
}
