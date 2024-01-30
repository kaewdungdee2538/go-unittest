//go:build integration

package handlers_test

import (
	"fmt"
	"gotest/handlers"
	"gotest/repositories"
	"gotest/services"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalcualteDiscountIntegrationService(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		amount := 100
		expected := 80

		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{
			ID:              1,
			PurchaseMin:     100,
			DiscountPercent: 20,
		}, nil)

		promoService := services.NewPromotionService(promoRepo)
		promoHandler := handlers.NewPromotionHanlder(promoService)

		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		//http://localhost:8000/calculate?amount=100
		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)
		//Act
		response, _ := app.Test(req)

		defer response.Body.Close()

		//Assert
		if assert.Equal(t, fiber.StatusOK, response.StatusCode) {
			body, _ := io.ReadAll(response.Body)
			assert.Equal(t, strconv.Itoa(expected), string(body))
		}
	})

}
