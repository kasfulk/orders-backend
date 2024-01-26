package repository

import (
	"github.com/kasfulk/orders-backend/internal/utils"
	"github.com/kasfulk/orders-backend/services/domain"
	"github.com/kasfulk/orders-backend/services/models"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(DB *gorm.DB) domain.OrderRepository {
	return &orderRepository{DB}
}

func (o *orderRepository) FetchAll(c *fiber.Ctx, page int, size int) (pager paginate.Page, rowLen int64, err error) {
	var domainOrders []models.Order
	pg := paginate.New()

	result := o.DB.Find(&domainOrders)
	pager = pg.With(result).Request(c.Request()).Response(&domainOrders)
	utils.CleanupDBConnection(o.DB)
	return pager, result.RowsAffected, result.Error
}

func (o *orderRepository) FetchOneByID(c *fiber.Ctx, id uint) (res models.Order, err error) {
	result := o.DB.First(&res, id)
	utils.CleanupDBConnection(o.DB)
	return res, result.Error
}

func (o *orderRepository) Save(c *fiber.Ctx) (res models.Order, err error) {
	order := new(models.Order)
	if err = c.BodyParser(&order); err != nil {
		return
	}

	o.DB.Create(&order)
	utils.CleanupDBConnection(o.DB)
	return *order, nil
}

func (o *orderRepository) Edit(c *fiber.Ctx, id uint) (res models.Order, err error) {
	result := o.DB.First(&res, id)
	if result.Error != nil {
		return models.Order{}, result.Error
	}
	if err = c.BodyParser(&res); err != nil {
		return models.Order{}, err
	}
	o.DB.Where("id = ? ", id).Save(&res)
	utils.CleanupDBConnection(o.DB)
	return res, nil
}

func (o *orderRepository) SoftDelete(c *fiber.Ctx, id uint) (err error) {
	var order models.Order
	result := o.DB.Where("id = ?", id).Delete(&order)
	if result.Error != nil {
		return result.Error
	}

	utils.CleanupDBConnection(o.DB)
	return
}
