package b_badsolution

import "fmt"

// This solution uses extensive use of if else statements or switch cases. This solution is not extensible and
// maintainable.

type Canvas struct {
	tool Tool
}

func NewCanvas(tool Tool) *Canvas {
	return &Canvas{tool: tool}
}

func (c *Canvas) MouseDown() {
	switch c.tool {
	case Selection:
		fmt.Println("Selection cursor")
	case Brush:
		fmt.Println("Brush cursor")
	case Eraser:
		fmt.Println("Eraser cursor")
	}
}

func (c *Canvas) MouseUp() {
	switch c.tool {
	case Selection:
		fmt.Println("Draws a rectangular")
	case Brush:
		fmt.Println("Draws a line")
	case Eraser:
		fmt.Println("Erase something")
	}
}

func (c *Canvas) Tool() Tool {
	return c.tool
}

func (c *Canvas) SetTool(tool Tool) {
	c.tool = tool
}
