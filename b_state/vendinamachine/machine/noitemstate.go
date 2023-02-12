package machine

import "fmt"

type NoItem struct {
	vendingMachine *VendingMachine
}

func NewNoItem(vendingMachine *VendingMachine) *NoItem {
	return &NoItem{vendingMachine: vendingMachine}
}

func (ni *NoItem) RequestItem() error {
	return fmt.Errorf("no item available")
}

func (ni *NoItem) AddItem(amount int) error {
	ni.vendingMachine.itemCount += amount
	ni.vendingMachine.SetState(NewHasItem(ni.vendingMachine))
	return nil
}

func (ni *NoItem) InsertMoney(amount int) error {
	return fmt.Errorf("no item available")
}

func (ni *NoItem) DispenseItem() (string, error) {
	return "", fmt.Errorf("no item available")
}
