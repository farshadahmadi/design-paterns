package machine

import "fmt"

type HasMoney struct {
	vendingMachine *VendingMachine
}

func NewHasMoney(vendingMachine *VendingMachine) *HasMoney {
	return &HasMoney{vendingMachine: vendingMachine}
}

func (hm *HasMoney) RequestItem() error {
	return fmt.Errorf("item already requested")
}

func (hm *HasMoney) AddItem(amount int) error {
	hm.vendingMachine.itemCount += amount
	return nil
}

func (hm *HasMoney) InsertMoney(amount int) error {
	return fmt.Errorf("money already inserted")
}

func (hm *HasMoney) DispenseItem() (string, error) {
	item := "item"
	hm.vendingMachine.SetState(NewNoItem(hm.vendingMachine))
	return item, nil
}
