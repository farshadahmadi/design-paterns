package machine

type State interface {
	AddItem(amount int) error
	RequestItem() error
	InsertMoney(amount int) error
	DispenseItem() (string, error)
}
