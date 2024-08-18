package handler

import (
	"fmt"
	"rizkiwhy-dating-app/api/presenter"
	"rizkiwhy-dating-app/pkg/order"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type OrderHandlerImpl interface {
	PurchasePremiumPackage() fiber.Handler
}

type OrderHandler struct {
	OrderService *order.OrderService
}

func NewOrderHandler(orderService *order.OrderService) OrderHandlerImpl {
	return &OrderHandler{
		OrderService: orderService,
	}
}

func (h *OrderHandler) PurchasePremiumPackage() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		var request presenter.PurchasePremiumPackageRequest

		if err = c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(err))
		}

		request.UserID, err = uuid.Parse(c.Locals("user_id").(string))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(err))
		}

		if errs := validator.New().Struct(request); errs != nil {
			for _, vErr := range errs.(validator.ValidationErrors) {
				return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(fmt.Errorf("%s should be %s", vErr.Field(), vErr.Tag())))
			}
		}

		if err = h.OrderService.PurchasePremiumPackage(request); err != nil {
			if er, ok := err.(*fiber.Error); ok {
				return c.Status(er.Code).JSON(presenter.GlobalErrorResponse(err))
			}

			return c.Status(fiber.StatusInternalServerError).JSON(presenter.GlobalErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(presenter.GlobalSuccessResponse())
	}
}
