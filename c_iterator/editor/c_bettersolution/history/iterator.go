package history

import (
	"github.com/farshadahmadi/c_iterator/editor/a_problem/editorandstate"
)

type EditorHistoryIterator struct {
	editorHistory *EditorStateHistory
	index         int
}

func NewEditorHistoryIterator(editorHistory *EditorStateHistory) *EditorHistoryIterator {
	return &EditorHistoryIterator{editorHistory: editorHistory, index: 0}
}

func (ehi *EditorHistoryIterator) Next() (editorandstate.EditorState, bool) {
	if ehi.index < len(ehi.editorHistory.history) {
		state := ehi.editorHistory.history[ehi.index]
		ehi.index++
		return state, true
	}

	return nil, false
}
