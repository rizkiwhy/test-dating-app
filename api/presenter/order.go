package presenter

import (
	"github.com/google/uuid"
)

type PurchasePremiumPackageRequest struct {
	UserID             uuid.UUID `json:"userID" validate:"required"`
	PremiumPackageCode string    `json:"premiumPackageCode" validate:"required"`
}
