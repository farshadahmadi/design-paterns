package main

import (
	"fmt"

	"github.com/farshadahmadi/memento/editor/actualsolution"
	"github.com/farshadahmadi/memento/editor/badsolution"
	"github.com/farshadahmadi/memento/editor/bettersolution"
)

func main() {
	// client code using bad solution
	e1 := badsolution.NewEditor("a")
	e1.SetContent("b")
	e1.SetContent("c")
	fmt.Println(e1.Content())
	e1.Undo()
	fmt.Println(e1.Content())
	e1.Undo()
	fmt.Println(e1.Content())
	fmt.Println("---------")

	// client code using better solution
	// This solution delegate all history management stuff to Editor service. This violates Single Responsibility
	// principle as Editor service ia also doing history management! However, client code is cleaner as history
	// management is abstracted away. Compare this solution to the next one where client handles history management.
	e2 := bettersolution.NewEditor()
	e2.SetContent("a")
	e2.SetContent("b")
	e2.SetContent("c")
	fmt.Println(e2.Content())
	e2.Undo()
	fmt.Println(e2.Content())
	e2.Undo()
	fmt.Println(e2.Content())
	fmt.Println("---------")

	// client code using actual solution
	// This solution allows client code to handle history management. Thus, Editor service does not need to handle history management
	// conforming to Single Responsibility Principle.
	esh := actualsolution.NewEditorStateHistory()
	e3 := actualsolution.NewEditor()
	e3.SetContent("a")
	esh.PushState(e3.GetState())

	e3.SetContent("b")
	esh.PushState(e3.GetState())

	e3.SetContent("c")
	esh.PushState(e3.GetState())

	es := esh.PopState()
	e3.SetState(es)
	fmt.Println(e3.Content())

	es = esh.PopState()
	e3.SetState(es)
	fmt.Println(e3.Content())

	es = esh.PopState()
	e3.SetState(es)
	fmt.Println(e3.Content())
	fmt.Println("---------")
}
