package presenter

import "github.com/gofiber/fiber/v2"

func GlobalErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}

func GlobalSuccessResponse() *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   "",
		"error":  nil,
	}
}
