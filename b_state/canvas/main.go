package main

import (
	"github.com/farshadahmadi/b_state/canvas/b_badsolution"
	"github.com/farshadahmadi/b_state/canvas/c_bettersolution"
)

func main() {
	// Client code using a bad solution

	canvas := b_badsolution.NewCanvas(b_badsolution.Selection)
	canvas.MouseDown()
	canvas.MouseUp()

	canvas.SetTool(b_badsolution.Brush)
	canvas.MouseDown()
	canvas.MouseUp()

	canvas.SetTool(b_badsolution.Eraser)
	canvas.MouseDown()
	canvas.MouseUp()

	// Client code using better solution
	canvas1 := c_bettersolution.NewCanvas(c_bettersolution.NewSelection())
	canvas1.MouseDown()
	canvas1.MouseUp()

	canvas1.SetTool(c_bettersolution.NewEraser())
	canvas1.MouseDown()
	canvas1.MouseUp()

	canvas1.SetTool(c_bettersolution.NewBrush())
	canvas1.MouseDown()
	canvas1.MouseUp()
}
