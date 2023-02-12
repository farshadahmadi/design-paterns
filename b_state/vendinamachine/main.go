package main

import (
	"fmt"

	"github.com/farshadahmadi/b_state/vendingmachine/machine"
)

func main() {
	vm := machine.NewVendingMachine()
	err := vm.RequestItem()
	if err != nil {
		fmt.Println(err)
	}

	err = vm.AddItem(1)
	if err != nil {
		fmt.Println(err)
	}

	err = vm.RequestItem()
	if err != nil {
		fmt.Println(err)
	}

	err = vm.InsertMoney(5)
	if err != nil {
		fmt.Println(err)
	}

	err = vm.InsertMoney(10)
	if err != nil {
		fmt.Println(err)
	}

	item, err := vm.DispenseItem()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(item)
}
