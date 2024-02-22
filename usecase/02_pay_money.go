package usecase

import "hiromi-mitsuoka/vending_machine/domain"

// お題2
func DoPayMoney(money int) (*domain.Beverage, error) {
	vm := domain.VendingMachine{}
	if err := vm.Insert(money); err != nil {
		return nil, err
	}

	return vm.PushButton("Coke"), nil
}
