package domain

import "fmt"

type VendingMachineAction interface {
	PushButton(buttonName string) *Beverage
	Despence(name string) *Beverage
	Insert(money int) error
}

type (
	VendingMachine struct {
		VendingMachineAction
	}
)

// func NewVendingMachine(action VendingMachineAction) *VendingMachine {
// 	return &VendingMachine{
// 		VendingMachineAction: action,
// 	}
// }

func (vm VendingMachine) PushButton(buttonName string) *Beverage {
	return vm.Despence(buttonName)
}

func (vm VendingMachine) Despence(name string) *Beverage {
	return &Beverage{
		Name:           BeverageName(name),
		BeverageAction: nil,
	}
}

func (vm VendingMachine) Insert(money int) error {
	if money != 100 {
		return fmt.Errorf("money is not 100")
	}

	return nil
}
