package router

import (
	"rizkiwhy-dating-app/api/handler"
	"rizkiwhy-dating-app/middleware"

	"github.com/gofiber/fiber/v2"
)

func OrderRouter(app fiber.Router, h *handler.OrderHandler) {
	orderAPI := app.Group("/orders")
	orderAPI.Post("/purchase-premium-package", middleware.Protected(), h.PurchasePremiumPackage())
}
