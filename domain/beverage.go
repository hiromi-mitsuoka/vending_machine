package domain

type BeverageAction interface{}

type (
	Beverage struct {
		Name BeverageName
		BeverageAction
	}

	BeverageName string
)

// func NewBeverage(name string, action BeverageAction) *Beverage {
// 	return &Beverage{
// 		Name:           BeverageName(name),
// 		BeverageAction: action,
// 	}
// }
