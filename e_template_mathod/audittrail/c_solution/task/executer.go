package task

import (
	"github.com/farshadahmadi/e_template_method/audittrail/c_solution/audittrail"
)

type Executer struct {
	auditTrail *audittrail.AuditTrail
}

func NewExecuter(auditTrail *audittrail.AuditTrail) *Executer {
	return &Executer{auditTrail: auditTrail}
}

func (e *Executer) Execute(task Task) {
	e.auditTrail.Record()
	task.execute()
}
