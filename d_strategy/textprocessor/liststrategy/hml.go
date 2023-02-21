package liststrategy

import (
	"fmt"
	"strings"
)

type HtmlListStrategy struct {
}

func NewHtmlListStrategy() *HtmlListStrategy {
	return &HtmlListStrategy{}
}

func (hls *HtmlListStrategy) Start(builder *strings.Builder) {
	builder.WriteString("<ul>\n")
}

func (hls *HtmlListStrategy) End(builder *strings.Builder) {
	builder.WriteString("</ul>\n")
}

func (hls *HtmlListStrategy) Add(builder *strings.Builder, item string) {
	builder.WriteString(fmt.Sprintf("  <li>%s</li>\n", item))
}
