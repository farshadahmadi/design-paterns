package main

import (
	"fmt"

	"github.com/farshadahmadi/d_strategy/textprocessor/liststrategy"
	"github.com/farshadahmadi/d_strategy/textprocessor/processor"
)

func main() {
	tp := processor.NewTextProcessor(liststrategy.NewMarkDownListStrategy())
	tp.ConvertToListItems([]string{"item 1", "item 2", "item 3"})
	fmt.Println(tp)

	tp.Reset()
	tp.SetListStrategy(liststrategy.NewHtmlListStrategy())
	tp.ConvertToListItems([]string{"item 1", "item 2", "item 3"})
	fmt.Println(tp)

}
