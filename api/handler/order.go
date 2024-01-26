package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kasfulk/orders-backend/internal/order-system/delivery"
	"github.com/kasfulk/orders-backend/services/domain"
)

func NewOrderHandler(c fiber.Router, uc domain.OrderUsecase) {
	handler := &delivery.OrderHandler{
		OrderUsecase: uc,
	}
	g := c.Group("/orders")
	g.Get("", handler.FetchAll)
	g.Get("/:id", handler.FetchOneByID)
	g.Post("", handler.Save)
	g.Put("/:id", handler.Edit)
	g.Delete("/:id", handler.SoftDelete)
}
