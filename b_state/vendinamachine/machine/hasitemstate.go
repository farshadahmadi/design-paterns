package machine

import "fmt"

type HasItem struct {
	vendingMachine *VendingMachine
}

func NewHasItem(vendingMachine *VendingMachine) *HasItem {
	return &HasItem{vendingMachine: vendingMachine}
}

func (hi *HasItem) RequestItem() error {
	hi.vendingMachine.SetState(NewItemRequested(hi.vendingMachine))
	return nil
}

func (hi *HasItem) AddItem(amount int) error {
	hi.vendingMachine.itemCount += amount
	return nil
}

func (hi *HasItem) InsertMoney(amount int) error {
	if amount < hi.vendingMachine.itemCount*hi.vendingMachine.itemPrice {
		return fmt.Errorf("not enough money")
	}

	hi.vendingMachine.SetState(NewHasMoney(hi.vendingMachine))
	return nil
}

func (hi *HasItem) DispenseItem() (string, error) {
	return "", fmt.Errorf("no money inserted")
}
