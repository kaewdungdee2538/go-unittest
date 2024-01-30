package handlers_test

import (
	"fmt"
	"gotest/handlers"
	"gotest/services"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//Arrage
		amount := 100
		expected := 80

		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(expected, nil)

		promoHandler := handlers.NewPromotionHanlder(promoService)

		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		//http://localhost:8000/calculate?amount=100
		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)
		//Act
		response, _ := app.Test(req)

		defer response.Body.Close()

		//Assert
		if assert.Equal(t, fiber.StatusOK, response.StatusCode){
			body, _ := io.ReadAll(response.Body)
			assert.Equal(t, strconv.Itoa(expected), string(body))
		}

	})
}
