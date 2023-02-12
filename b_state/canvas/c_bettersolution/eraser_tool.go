package c_bettersolution

import "fmt"

type Eraser struct {
}

func NewEraser() *Eraser {
	return &Eraser{}
}

func (e Eraser) MouseUp() {
	fmt.Println("Erase something")
}

func (e Eraser) MouseDown() {
	fmt.Println("Eraser cursor")
}
