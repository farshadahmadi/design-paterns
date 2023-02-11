package main

import (
	"fmt"

	"github.com/farshadahmadi/a_memento/document/document"
	"github.com/farshadahmadi/a_memento/document/history"
)

func main() {
	dh := history.NewDocumentHistory()
	d := document.NewDocument("Hello", "font name 1", "font size 1")
	dh.PushState(d.CreateState())

	d.SetFotName("font name 2")
	dh.PushState(d.CreateState())

	d.SetFontSize("font size 2")
	dh.PushState(d.CreateState())

	d.SetContent("World!")
	d.LoadState(dh.PopState())
	fmt.Printf("%+v\n", d)

	d.LoadState(dh.PopState())
	fmt.Printf("%+v\n", d)

	d.LoadState(dh.PopState())
	fmt.Printf("%+v\n", d)
}
