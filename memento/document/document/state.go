package document

type State struct {
	content  string
	fotName  string
	fontSize string
}

func newState(content string, fotName string, fontSize string) *State {
	return &State{content: content, fotName: fotName, fontSize: fontSize}
}

func (s State) Content() string {
	return s.content
}

func (s State) FotName() string {
	return s.fotName
}

func (s State) FontSize() string {
	return s.fontSize
}
