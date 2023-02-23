package task

import (
	"fmt"

	"github.com/farshadahmadi/e_template_method/audittrail/b_bad_solution/audittrail"
)

type GenerateReport struct {
	auditTrail *audittrail.AuditTrail
}

func NewGenerateReport(auditTrail *audittrail.AuditTrail) *GenerateReport {
	return &GenerateReport{auditTrail: auditTrail}
}

func (gr *GenerateReport) Execute() {
	gr.auditTrail.Record()
	fmt.Println("Generate report task ...")
}
