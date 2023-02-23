package main

import (
	"fmt"

	"github.com/farshadahmadi/e_template_method/audittrail/b_bad_solution/audittrail"
	"github.com/farshadahmadi/e_template_method/audittrail/b_bad_solution/task"
	audittrail2 "github.com/farshadahmadi/e_template_method/audittrail/c_solution/audittrail"
	task2 "github.com/farshadahmadi/e_template_method/audittrail/c_solution/task"
)

// Some tasks need to be done. Our problem is that we want to do a repetitive sub task as part of all these tasks.
// For example, our tasks are transferring money and generating report. But we want to audit trail them. Meaning whenever
// a money transfer is done, it is recorded by an audit tail. Th same goes to generating report. So we have info who did what.
// Think about a solution to do that!

func main() {
	// Client cod for The trivial bas solution.
	// Although single responsibility principle is followed here by putting codes related to transferring money and
	// generating reports into their own classes class, but there are two issues:
	// - DRY is not followed. You can see the same code in GenerateReport class for audit trail related code
	// - if later somebody add a task, no body is forcing them to record and audit trail!
	at := audittrail.NewAuditTrail()
	tm := task.NewTransferMoney(at)
	tm.Execute()
	gr := task.NewGenerateReport(at)
	gr.Execute()

	fmt.Println("=================")
	// Go does not have inheritance. So conventional template method design pattern which is based on inheritance
	// is not possible. But as we know, most inheritance-based approaches can be implemented using composition/delegation;
	// This will return us back to strategy pattern.
	at1 := audittrail2.NewAuditTrail()
	te := task2.NewExecuter(at1)
	te.Execute(task2.NewMoneyTransferTask())
	te.Execute(task2.NewGenerateReportTask())
}
