package main

import (
	"fmt"

	"github.com/farshadahmadi/c_iterator/editor/b_badsolution/editorandstate"
	"github.com/farshadahmadi/c_iterator/editor/b_badsolution/history"

	bettereditorandstate "github.com/farshadahmadi/c_iterator/editor/c_bettersolution/editorandstate"
	betterhistory "github.com/farshadahmadi/c_iterator/editor/c_bettersolution/history"
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

	fmt.Println("--------------")

	// Proper solution is to define an iterator interface. There will a concrete implementation of the iterator interface
	// specific for history service, lets call it EditorHistoryIterator. The iterator concrete implement is a separate service
	// to follow Single Responsibility Principle, it is responsible only for iterating the mementos. It needs a reference
	// to the history class to access mementos. History class provides client with a exported function to create a fresh iterator.
	// The return iterator is of type iterator interface, this is programming to an interface!

	esh1 := betterhistory.NewEditorStateHistory()
	e1 := bettereditorandstate.NewEditor()
	e1.SetContent("a")
	esh1.PushState(e1.GetState())

	e1.SetContent("b")
	esh1.PushState(e1.GetState())

	e1.SetContent("c")
	esh1.PushState(e1.GetState())

	iterator := esh1.CreateIterator()

	for state, ok := iterator.Next(); ok; state, ok = iterator.Next() {
		fmt.Println(state)
	}
}
