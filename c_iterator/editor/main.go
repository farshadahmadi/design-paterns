package main

import (
	"fmt"

	"github.com/farshadahmadi/c_iterator/editor/a_problem/editorandstate"
	"github.com/farshadahmadi/c_iterator/editor/a_problem/history"
)

func main() {
	esh := history.NewEditorStateHistory()
	e := editorandstate.NewEditor()
	e.SetContent("a")
	esh.PushState(e.GetState())

	e.SetContent("b")
	esh.PushState(e.GetState())

	e.SetContent("c")
	esh.PushState(e.GetState())

	for state, ok := esh.Next(); ok; state, ok = esh.Next() {
		fmt.Println(state)
	}

	esh.Reset()
	for state, ok := esh.Next(); ok; state, ok = esh.Next() {
		fmt.Println(state)
	}
}
