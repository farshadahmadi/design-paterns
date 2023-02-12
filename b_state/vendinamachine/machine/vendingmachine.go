package machine

type VendingMachine struct {
	state     State
	itemCount int
	itemPrice int
}

func NewVendingMachine() *VendingMachine {
	v := &VendingMachine{
		itemCount: 0,
		itemPrice: 10,
	}
	v.state = NewNoItem(v)
	return v
}

func (v *VendingMachine) RequestItem() error {
	return v.state.RequestItem()
}

func (v *VendingMachine) AddItem(amount int) error {
	return v.state.AddItem(amount)
}

func (v *VendingMachine) InsertMoney(amount int) error {
	return v.state.InsertMoney(amount)
}

func (v *VendingMachine) DispenseItem() (string, error) {
	return v.state.DispenseItem()
}

func (v *VendingMachine) SetState(state State) {
	v.state = state
}
