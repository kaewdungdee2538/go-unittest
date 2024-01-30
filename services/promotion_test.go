package services_test

import (
	"errors"
	"fmt"
	"gotest/repositories"
	"gotest/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateDiscount(t *testing.T) {
	type testCase struct {
		name            string
		puchaseMin      int
		discountPercent int
		amount          int
		expected        int
	}
	cases := []testCase{
		{name: "applied 100", puchaseMin: 100, discountPercent: 20, amount: 100, expected: 80},
		{name: "applied 200", puchaseMin: 200, discountPercent: 20, amount: 200, expected: 160},
		{name: "applied 300", puchaseMin: 300, discountPercent: 20, amount: 300, expected: 240},
		{name: "not applied 50", puchaseMin: 50, discountPercent: 20, amount: 50, expected: 50},
	}
	promoRepo := repositories.NewPromotionRepositoryMock()

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			//Arrage -> mock data
			promoRepo.On("GetPromotion").Return(repositories.Promotion{
				ID:              1,
				PurchaseMin:     c.puchaseMin,
				DiscountPercent: c.discountPercent,
			}, nil)

			// Act -> doing
			promoService := services.NewPromotionService(promoRepo)
			discount, _ := promoService.CalculateDiscount(c.amount)

			// Assert validate
			assert.Equal(t, c.expected, discount)
		})
	}

	// test case zero
	t.Run("applied zero value", func(t *testing.T) {
		promoRepo := repositories.NewPromotionRepositoryMock()
		//Arrage -> mock data
		promoRepo.On("GetPromotion").Return(repositories.Promotion{
			ID:              1,
			PurchaseMin:     100,
			DiscountPercent: 20,
		}, nil)

		// Act -> doing
		promoService := services.NewPromotionService(promoRepo)
		_, err := promoService.CalculateDiscount(0)
		fmt.Println(err.Error())
		// Assert validate
		assert.ErrorIs(t, err, services.ErrZeroAmount)
		// not call function GetPromotion when function is next
		promoRepo.AssertNotCalled(t, "GetPromotion")
	})

	// test case repo error
	t.Run("applied repo error", func(t *testing.T) {
		promoRepo := repositories.NewPromotionRepositoryMock()
		//Arrage -> mock data
		promoRepo.On("GetPromotion").Return(repositories.Promotion{}, errors.New("repository error"))

		// Act -> doing
		promoService := services.NewPromotionService(promoRepo)
		_, err := promoService.CalculateDiscount(100)
		fmt.Println(err.Error())
		// Assert validate
		assert.ErrorIs(t, err, services.ErrRepository)
	})
}
