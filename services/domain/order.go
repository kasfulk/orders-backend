package domain

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kasfulk/orders-backend/services/models"
	"github.com/morkid/paginate"
)

type OrderUsecase interface {
	FetchAll(c *fiber.Ctx, page int, size int) (paginate.Page, int64, error)
	FetchOneByID(c *fiber.Ctx) (models.Order, error)
	Save(c *fiber.Ctx) (models.Order, error)
	Edit(c *fiber.Ctx) (models.Order, error)
	SoftDelete(c *fiber.Ctx) error
}

type OrderRepository interface {
	FetchAll(c *fiber.Ctx, page int, size int) (paginate.Page, int64, error)
	FetchOneByID(c *fiber.Ctx, id uint) (models.Order, error)
	Save(c *fiber.Ctx) (models.Order, error)
	Edit(c *fiber.Ctx, id uint) (models.Order, error)
	SoftDelete(c *fiber.Ctx, id uint) error
}
