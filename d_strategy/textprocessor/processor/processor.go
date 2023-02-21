package processor

import (
	"strings"

	"github.com/farshadahmadi/d_strategy/textprocessor/liststrategy"
)

type TextProcessor struct {
	builder      *strings.Builder
	listStrategy liststrategy.ListStrategy[string]
}

func NewTextProcessor(listStrategy liststrategy.ListStrategy[string]) *TextProcessor {
	return &TextProcessor{
		listStrategy: listStrategy,
		builder:      &strings.Builder{},
	}
}

func (tp *TextProcessor) Reset() {
	tp.builder.Reset()
}

func (tp *TextProcessor) ConvertToListItems(items []string) {
	tp.listStrategy.Start(tp.builder)
	for _, item := range items {
		tp.listStrategy.Add(tp.builder, item)
	}
	tp.listStrategy.End(tp.builder)
}

func (tp *TextProcessor) String() string {
	return tp.builder.String()
}

func (tp *TextProcessor) SetListStrategy(listStrategy liststrategy.ListStrategy[string]) {
	tp.listStrategy = listStrategy
}
