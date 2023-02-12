package c_bettersolution

import "fmt"

type Selection struct {
}

func NewSelection() *Selection {
	return &Selection{}
}

func (s Selection) MouseUp() {
	fmt.Println("Draws a rectangular")
}

func (s Selection) MouseDown() {
	fmt.Println("Selection cursor")
}
