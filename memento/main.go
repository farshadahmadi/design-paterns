package main

import (
	"fmt"

	"github.com/farshadahmadi/memento/editorsolution"
)

func main() {
	// client code using bad solution
	//e := editorbadsolution.NewEditor("a")
	//e.SetContent("b")
	//e.SetContent("c")
	//fmt.Println(e.Content())
	//e.Undo()
	//fmt.Println(e.Content())
	//e.Undo()
	//fmt.Println(e.Content())

	// client code using right solution
	ed := editorsolution.NewEditor()
	ed.SetContent("a")
	ed.SetContent("b")
	fmt.Println(ed.Content())
	ed.Undo()
	fmt.Println(ed.Content())
	ed.SetContent("c")
	ed.Undo()
	fmt.Println(ed.Content())

}
