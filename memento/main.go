package main

import (
	"fmt"

	"github.com/farshadahmadi/memento/editor"
)

func main() {
	fmt.Println("Test")
	e := editor.NewEditor("a")
	e.SetContent("b")

}
