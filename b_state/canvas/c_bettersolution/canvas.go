package c_bettersolution

type Canvas struct {
	tool Tool
}

func NewCanvas(tool Tool) *Canvas {
	return &Canvas{tool: tool}
}

func (c *Canvas) MouseDown() {
	c.tool.MouseDown()
}

func (c *Canvas) MouseUp() {
	c.tool.MouseUp()
}

func (c *Canvas) Tool() Tool {
	return c.tool
}

func (c *Canvas) SetTool(tool Tool) {
	c.tool = tool
}
