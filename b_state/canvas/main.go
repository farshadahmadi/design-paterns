package main

import (
	"github.com/farshadahmadi/b_state/canvas/b_badsolution"
)

func main() {
	// Client code using a bad solution

	canvas := b_badsolution.NewCanvas(b_badsolution.Selection)
	canvas.MouseDown()
	canvas.MouseUp()

	canvas.SetTool(b_badsolution.Brush)
	canvas.MouseDown()
	canvas.MouseUp()

	canvas.SetTool(b_badsolution.Brush)
	canvas.MouseDown()
	canvas.MouseUp()

	canvas.SetTool(b_badsolution.Eraser)
	canvas.MouseDown()
	canvas.MouseUp()
}
