package main

import (
	"fmt"

	"github.com/farshadahmadi/a_memento/editor/b_badsolution"
	"github.com/farshadahmadi/a_memento/editor/c_bettersolution"
	"github.com/farshadahmadi/a_memento/editor/d_actualsolution"
	"github.com/farshadahmadi/a_memento/editor/e_bestsolution/editorandstate"
	"github.com/farshadahmadi/a_memento/editor/e_bestsolution/history"
)

func main() {
	// client code using bad solution
	e1 := b_badsolution.NewEditor("a")
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
	e2 := c_bettersolution.NewEditor()
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
	esh := d_actualsolution.NewEditorStateHistory()
	e3 := d_actualsolution.NewEditor()
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

	// client code using best solution.
	// This solution is almost like previous solution with a difference, now editor and editor state are in one package,
	// letting editor to access all unexported fields of editor state. However history which is in another package only
	// has access to the exported (actually interface) methods. This is like using nested classes in Java which allows
	// parent class (Editor) to all private fields of nested class (State), but encapsulates State from other classes (
	// History). In Go, we can achieve the same same by putting structs into one package. Using interface is not
	// necessary as History class can access only exported methods, but it is recommended!
	esh1 := history.NewEditorStateHistory()
	e4 := editorandstate.NewEditor()
	e4.SetContent("a")
	esh1.PushState(e4.GetState())

	e4.SetContent("b")
	esh1.PushState(e4.GetState())

	e4.SetContent("c")
	esh1.PushState(e4.GetState())

	es1 := esh1.PopState()
	e4.SetState(es1)
	fmt.Println(e4.Content())

	es1 = esh1.PopState()
	e4.SetState(es1)
	fmt.Println(e4.Content())

	es1 = esh1.PopState()
	e4.SetState(es1)
	fmt.Println(e4.Content())
	fmt.Println("---------")
}
