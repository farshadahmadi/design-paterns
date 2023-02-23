package audittrail

import "fmt"

type AuditTrail struct {
}

func NewAuditTrail() *AuditTrail {
	return &AuditTrail{}
}
func (at AuditTrail) Record() {
	fmt.Println("Audit recorded ...")
}
