package history

import "github.com/farshadahmadi/memento/document/document"

type DocumentHistory struct {
	history []*document.State
}

func NewDocumentHistory() *DocumentHistory {
	return &DocumentHistory{
		history: make([]*document.State, 0),
	}
}

func (dh *DocumentHistory) PushState(s *document.State) {
	dh.history = append(dh.history, s)
}

func (dh *DocumentHistory) PopState() *document.State {
	s := dh.history[len(dh.history)-1]
	dh.history = dh.history[:len(dh.history)-1]
	return s
}
