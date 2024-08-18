package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterUserRequest struct {
	Username        string `json:"username" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
	Name            string `json:"name" validate:"required"`
	Birthday        string `json:"birthday" validate:"required"`
	Gender          string `json:"gender" validate:"required,oneof=male female"`
	IsGenderVisible bool   `json:"isGenderVisible" validate:"required"`
}

func (request RegisterUserRequest) HashPassword() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	return string(bytes), err
}

func UserLoginSuccessResponse(t string) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   t,
		"error":  nil,
	}
}

func SearchPartnerProfileSuccessResponse(response SearchPartnerProfileResponse) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   response,
		"error":  nil,
	}
}

type SearchPartnerProfileRequest struct {
	UserID uuid.UUID `json:"userID" validate:"required"`
}

type SearchPartnerProfileResponse struct {
	PartnerID          uuid.UUID `json:"partnerID"`
	Age                string    `json:"age"`
	Gender             *string   `json:"gender"`
	IsVerifiedAccount  bool      `json:"isVerifiedAccount"`
	Interested         string    `json:"interested"`
	RelationshipType   string    `json:"relationshipType"`
	RelationshipTypeID int
}

type ImpressPartnerProfileRequest struct {
	Sender   uuid.UUID `json:"sender" validate:"required"`
	Receiver uuid.UUID `json:"receiver" validate:"required"`
	Swipe    string    `json:"swipe" validate:"required,oneof=pass like"`
}
