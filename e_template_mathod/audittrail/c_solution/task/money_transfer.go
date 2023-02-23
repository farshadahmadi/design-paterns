package task

import "fmt"

type MoneyTransferTask struct{}

func NewMoneyTransferTask() *MoneyTransferTask {
	return &MoneyTransferTask{}
}

func (mts *MoneyTransferTask) execute() {
	fmt.Println("Money transfer task ...")
}
