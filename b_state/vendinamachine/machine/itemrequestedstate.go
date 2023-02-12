package machine

import "fmt"

type ItemRequested struct {
	vendingMachine *VendingMachine
}

func NewItemRequested(vendingMachine *VendingMachine) *ItemRequested {
	return &ItemRequested{vendingMachine: vendingMachine}
}

func (ir *ItemRequested) RequestItem() error {
	return fmt.Errorf("item already requested money")
}

func (ir *ItemRequested) AddItem(amount int) error {
	ir.vendingMachine.itemCount += amount
	return nil
}

func (ir *ItemRequested) InsertMoney(amount int) error {
	if amount < ir.vendingMachine.itemCount*ir.vendingMachine.itemPrice {
		return fmt.Errorf("not enough money")
	}

	ir.vendingMachine.SetState(NewHasMoney(ir.vendingMachine))
	return nil
}

func (ir *ItemRequested) DispenseItem() (string, error) {
	return "", fmt.Errorf("not enough money")
}
