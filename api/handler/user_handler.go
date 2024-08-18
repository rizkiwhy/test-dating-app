package handler

import (
	"fmt"
	"rizkiwhy-dating-app/api/presenter"
	"rizkiwhy-dating-app/pkg/user"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandlerImpl interface {
	RegisterUser() fiber.Handler
	Login() fiber.Handler
	SearchPartnerProfile() fiber.Handler
}

type UserHandler struct {
	UserService *user.UserService
}

func NewUserHandler(userService *user.UserService) UserHandlerImpl {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) RegisterUser() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		var request presenter.RegisterUserRequest

		if err = c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(err))
		}

		if errs := validator.New().Struct(request); errs != nil {
			for _, vErr := range errs.(validator.ValidationErrors) {
				switch vErr.Tag() {
				case "oneof":
					err = fmt.Errorf("field '%s' must be one of the following values: %s", vErr.Field(), vErr.Param())
				case "required":
					err = fmt.Errorf("field '%s' is required", vErr.Field())
				default:
					err = fmt.Errorf("field '%s' validation failed: %s", vErr.Field(), vErr.Tag())
				}
			}

			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(err))
			}
		}

		if err = h.UserService.RegisterUser(request); err != nil {
			if er, ok := err.(*fiber.Error); ok {
				return c.Status(er.Code).JSON(presenter.GlobalErrorResponse(err))
			}

			return c.Status(fiber.StatusInternalServerError).JSON(presenter.GlobalErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(presenter.GlobalSuccessResponse())
	}
}

func (h *UserHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		var request presenter.LoginUserRequest
		if err = c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(err))
		}

		if errs := validator.New().Struct(request); errs != nil {
			for _, vErr := range errs.(validator.ValidationErrors) {
				switch vErr.Tag() {
				case "oneof":
					err = fmt.Errorf("field '%s' must be one of the following values: %s", vErr.Field(), vErr.Param())
				case "required":
					err = fmt.Errorf("field '%s' is required", vErr.Field())
				default:
					err = fmt.Errorf("field '%s' validation failed: %s", vErr.Field(), vErr.Tag())
				}
			}

			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(err))
			}
		}

		t, err := h.UserService.LoginUser(request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.GlobalErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(presenter.UserLoginSuccessResponse(t))
	}
}

func (h *UserHandler) SearchPartnerProfile() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		var request presenter.SearchPartnerProfileRequest

		request.UserID, err = uuid.Parse(c.Locals("user_id").(string))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(err))
		}

		if errs := validator.New().Struct(request); errs != nil {
			for _, vErr := range errs.(validator.ValidationErrors) {
				switch vErr.Tag() {
				case "oneof":
					err = fmt.Errorf("field '%s' must be one of the following values: %s", vErr.Field(), vErr.Param())
				case "required":
					err = fmt.Errorf("field '%s' is required", vErr.Field())
				default:
					err = fmt.Errorf("field '%s' validation failed: %s", vErr.Field(), vErr.Tag())
				}
			}

			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(err))
			}
		}

		partner, err := h.UserService.SearchPartnerProfile(request)
		if err != nil {
			if er, ok := err.(*fiber.Error); ok {
				return c.Status(er.Code).JSON(presenter.GlobalErrorResponse(err))
			}

			return c.Status(fiber.StatusInternalServerError).JSON(presenter.GlobalErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(presenter.SearchPartnerProfileSuccessResponse(*partner))
	}
}

func (h *UserHandler) ImpressPartnerProfile() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		var request presenter.ImpressPartnerProfileRequest

		request.Sender, err = uuid.Parse(c.Locals("user_id").(string))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(err))
		}

		request.Receiver, err = uuid.Parse(c.Params("partnerID"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(err))
		}

		request.Swipe = c.Params("swipe")
		if errs := validator.New().Struct(request); errs != nil {
			for _, vErr := range errs.(validator.ValidationErrors) {
				switch vErr.Tag() {
				case "oneof":
					err = fmt.Errorf("field '%s' must be one of the following values: %s", vErr.Field(), vErr.Param())
				case "required":
					err = fmt.Errorf("field '%s' is required", vErr.Field())
				default:
					err = fmt.Errorf("field '%s' validation failed: %s", vErr.Field(), vErr.Tag())
				}
			}

			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(err))
			}
		}

		err = h.UserService.ImpressPartnerProfile(request)
		if err != nil {
			if er, ok := err.(*fiber.Error); ok {
				return c.Status(er.Code).JSON(presenter.GlobalErrorResponse(err))
			}

			return c.Status(fiber.StatusInternalServerError).JSON(presenter.GlobalErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(presenter.GlobalSuccessResponse())
	}
}
