package task

import "fmt"

type GenerateReportTask struct{}

func NewGenerateReportTask() *GenerateReportTask {
	return &GenerateReportTask{}
}

func (grt *GenerateReportTask) execute() {
	fmt.Println("Generate report task ...")
}
