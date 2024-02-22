package usecase

import "hiromi-mitsuoka/vending_machine/domain"

// お題1
func DoPushButtonAndReturnCoke() *domain.Beverage {
	vm := domain.VendingMachine{}
	return vm.PushButton("Coke")
}
