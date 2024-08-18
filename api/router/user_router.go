package router

import (
	"rizkiwhy-dating-app/api/handler"
	"rizkiwhy-dating-app/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, h *handler.UserHandler) {
	userAPI := app.Group("/users")
	userAPI.Post("/register", h.RegisterUser())
	userAPI.Post("/login", h.Login())
	userAPI.Get("/search-partner-profile", middleware.Protected(), h.SearchPartnerProfile())
	userAPI.Post("/impress/:partnerID/:swipe", middleware.Protected(), h.ImpressPartnerProfile())
}
