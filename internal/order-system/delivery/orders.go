package delivery

import (
	"strconv"
	"strings"

	"github.com/kasfulk/orders-backend/api/response"
	"github.com/kasfulk/orders-backend/services/domain"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	OrderUsecase domain.OrderUsecase
}

func (o *OrderHandler) FetchAll(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return response.ReturnTheResponse(c, false, int(400), "Bad Request ?page= not found", nil)
	}
	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		return response.ReturnTheResponse(c, false, int(400), "Bad Request ?size= not found", nil)
	}
	listOrders, rowLen, err := o.OrderUsecase.FetchAll(c, page-1, size)
	if rowLen <= 0 {
		return response.ReturnTheResponse(c, true, int(404), "Record not Found", nil)
	}
	if err != nil {
		return response.ReturnTheResponse(c, true, int(404), err.Error(), nil)
	}
	return response.ReturnTheResponse(c, false, int(200), "", listOrders)
}

func (o *OrderHandler) FetchOneByID(c *fiber.Ctx) error {
	res, err := o.OrderUsecase.FetchOneByID(c)
	if err != nil {
		return response.ReturnTheResponse(c, true, int(404), err.Error(), nil)
	}
	return response.ReturnTheResponse(c, false, int(200), "", res)
}

func (o *OrderHandler) Save(c *fiber.Ctx) error {
	res, err := o.OrderUsecase.Save(c)
	if err != nil {
		return response.ReturnTheResponse(c, true, int(500), err.Error(), nil)
	}
	return response.ReturnTheResponse(c, false, int(200), "", res)
}

func (o *OrderHandler) Edit(c *fiber.Ctx) error {
	res, err := o.OrderUsecase.Edit(c)
	if err != nil {
		errMessage := err.Error()
		if strings.Contains(errMessage, "not found") {
			return response.ReturnTheResponse(c, true, int(404), err.Error(), nil)
		} else {
			return response.ReturnTheResponse(c, true, int(500), err.Error(), nil)
		}
	}
	return response.ReturnTheResponse(c, false, int(200), "", res)
}

func (o *OrderHandler) SoftDelete(c *fiber.Ctx) error {
	err := o.OrderUsecase.SoftDelete(c)
	if err != nil {
		return response.ReturnTheResponse(c, true, int(500), err.Error(), nil)
	}
	return response.ReturnTheResponse(c, false, int(200), "Deleted successfully", nil)
}
