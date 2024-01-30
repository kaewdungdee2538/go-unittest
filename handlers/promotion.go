package handlers

import (
	"fmt"
	"gotest/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PromotionHandlerHandler interface {
	CalculateDiscount(c *fiber.Ctx) error
}

type promotionHandler struct {
	promoService services.PromotionService
}

func NewPromotionHanlder(promoService services.PromotionService) promotionHandler {
	return promotionHandler{promoService}
}

func (h promotionHandler) CalculateDiscount(c *fiber.Ctx) error{
	//http://localhost:8000/calculate?amount=100

	amountStr := c.Query("amount")
	amount,err := strconv.Atoi(amountStr)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	discount,err := h.promoService.CalculateDiscount(amount)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.SendString(fmt.Sprint(discount))
}