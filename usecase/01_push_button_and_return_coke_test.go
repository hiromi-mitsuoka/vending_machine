package usecase_test

import (
	"hiromi-mitsuoka/vending_machine/domain"
	"hiromi-mitsuoka/vending_machine/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushButtonAndReturnCoke(t *testing.T) {
	t.Parallel()

	type (
		give struct{}
		want struct {
			b domain.Beverage
		}
	)

	tests := []struct {
		name string
		give give
		want want
	}{
		{
			name: "[OK] PushButtonAndReturnCoke() returns a coke beverage",
			give: give{},
			want: want{
				b: domain.Beverage{
					Name: "Coke",
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := usecase.DoPushButtonAndReturnCoke()
			assert.Equal(t, tt.want.b, *b)
		})
	}
}
