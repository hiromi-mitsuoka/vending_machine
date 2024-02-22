package usecase_test

import (
	"hiromi-mitsuoka/vending_machine/domain"
	"hiromi-mitsuoka/vending_machine/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPayMoney(t *testing.T) {
	t.Parallel()

	type (
		give struct {
			money int
		}
		want struct {
			b *domain.Beverage
		}
	)

	tests := []struct {
		name string
		give give
		want want
	}{
		{
			name: "[OK] PayMoney() returns a coke beverage",
			give: give{
				money: 100,
			},
			want: want{
				b: &domain.Beverage{
					Name: "Coke",
				},
			},
		},
		{
			name: "[NG] PayMoney() returns nil when money is less than 100",
			give: give{
				money: 99,
			},
			want: want{
				b: nil,
			},
		},
		{
			name: "[NG] PayMoney() returns nil when money is more than 100",
			give: give{
				money: 101,
			},
			want: want{
				b: nil,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b, err := usecase.DoPayMoney(tt.give.money)
			if err != nil {
				assert.Nil(t, b)
				assert.Error(t, err)
				return
			}
			assert.Equal(t, tt.want.b, b)
		})
	}
}
