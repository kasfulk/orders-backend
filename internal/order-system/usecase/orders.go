package usecase

import (
	"strconv"

	"github.com/kasfulk/orders-backend/services/domain"
	"github.com/kasfulk/orders-backend/services/models"

	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
)

type orderUsecase struct {
	orderRepo domain.OrderRepository
}

func NewOrderUsecase(r domain.OrderRepository) domain.OrderUsecase {
	return &orderUsecase{orderRepo: r}
}

func (o *orderUsecase) FetchAll(c *fiber.Ctx, page int, size int) (paginate.Page, int64, error) {
	paged, len, err := o.orderRepo.FetchAll(c, page, size)
	if err != nil {
		return paginate.Page{}, 0, err
	}
	return paged, len, nil
}

func (o *orderUsecase) FetchOneByID(c *fiber.Ctx) (res models.Order, err error) {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return
	}
	res, err = o.orderRepo.FetchOneByID(c, uint(id))
	return
}

func (o *orderUsecase) Save(c *fiber.Ctx) (res models.Order, err error) {
	res, err = o.orderRepo.Save(c)
	return
}

func (o *orderUsecase) Edit(c *fiber.Ctx) (res models.Order, err error) {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return
	}

	res, err = o.orderRepo.Edit(c, uint(id))
	return
}

func (o *orderUsecase) SoftDelete(c *fiber.Ctx) (err error) {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return
	}

	err = o.orderRepo.SoftDelete(c, uint(id))
	return
}
