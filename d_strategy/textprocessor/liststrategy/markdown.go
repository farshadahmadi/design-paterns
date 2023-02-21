package liststrategy

import (
	"fmt"
	"strings"
)

type MarkDownListStrategy struct{}

func NewMarkDownListStrategy() *MarkDownListStrategy {
	return &MarkDownListStrategy{}
}

func (m *MarkDownListStrategy) Start(builder *strings.Builder) {}

func (m *MarkDownListStrategy) End(builder *strings.Builder) {}

func (m *MarkDownListStrategy) Add(builder *strings.Builder, item string) {
	builder.WriteString(fmt.Sprintf("  * %s\n", item))
}
