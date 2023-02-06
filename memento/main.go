package main

import (
	"fmt"

	"github.com/farshadahmadi/memento/editorbadsolution"
)

func main() {
	// client code using bad solution
	e := editor.NewEditor("a")
	e.SetContent("b")
	e.SetContent("c")
	fmt.Println(e.Content())
	e.Undo()
	fmt.Println(e.Content())
	e.Undo()
	fmt.Println(e.Content())
}
