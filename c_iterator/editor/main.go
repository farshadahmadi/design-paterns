package main

import (
	"fmt"

	"github.com/farshadahmadi/c_iterator/editor/b_badsolution/editorandstate"
	"github.com/farshadahmadi/c_iterator/editor/b_badsolution/history"
)

func main() {
	// Trivial and bad solution is to introduce a getter for history slice. Now client can iterate over the mementos.
	// However, there are several issues with it. First, history object is exposed to client, violating encapsulation.
	// Second, if later internal data structure of history service (for keeping mementos) changes, then client code
	// breaks.

	esh := history.NewEditorStateHistory()
	e := editorandstate.NewEditor()
	e.SetContent("a")
	esh.PushState(e.GetState())

	e.SetContent("b")
	esh.PushState(e.GetState())

	e.SetContent("c")
	esh.PushState(e.GetState())

	for _, state := range esh.History() {
		fmt.Println(state)
	}
}
