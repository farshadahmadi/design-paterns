package task

import (
	"fmt"

	"github.com/farshadahmadi/e_template_method/audittrail/b_bad_solution/audittrail"
)

// Although single responsibility principle is followed here by putting codes related to transferring money nto its own
// class, but there are two issues:
// - DRY is not followed. You can see the same code in GenerateReport class for audit trail related code
// - if later somebody add a task, no body is forcing them to record and audit trail!

type TransferMoney struct {
	auditTrail *audittrail.AuditTrail
}

func NewTransferMoney(auditTrail *audittrail.AuditTrail) *TransferMoney {
	return &TransferMoney{auditTrail: auditTrail}
}

func (tm *TransferMoney) Execute() {
	tm.auditTrail.Record()
	fmt.Println("Money transfer task ...")
}
