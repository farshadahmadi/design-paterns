package c_bettersolution

import "fmt"

type Brush struct {
}

func NewBrush() *Brush {
	return &Brush{}
}

func (b Brush) MouseUp() {
	fmt.Println("Draws a line")
}

func (b Brush) MouseDown() {
	fmt.Println("Brush cursor")
}
