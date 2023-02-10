package document

type Document struct {
	content  string
	fotName  string
	fontSize string
}

func NewDocument(content string, fotName string, fontSize string) *Document {
	return &Document{content: content, fotName: fotName, fontSize: fontSize}
}

func (d *Document) Content() string {
	return d.content
}

func (d *Document) SetContent(content string) {
	d.content = content
}

func (d *Document) FotName() string {
	return d.fotName
}

func (d *Document) SetFotName(fotName string) {
	d.fotName = fotName
}

func (d *Document) FontSize() string {
	return d.fontSize
}

func (d *Document) SetFontSize(fontSize string) {
	d.fontSize = fontSize
}

func (d *Document) CreateState() *State {
	return newState(d.content, d.fotName, d.fontSize)
}

func (d *Document) LoadState(s *State) {
	d.content = s.content
	d.fotName = s.fotName
	d.fontSize = s.fontSize
}
