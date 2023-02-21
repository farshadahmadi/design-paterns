package liststrategy

import "strings"

type ListStrategy[T any] interface {
	Start(builder *strings.Builder)
	End(builder *strings.Builder)
	Add(builder *strings.Builder, item T)
}
