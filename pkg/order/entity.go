package order

import (
	"rizkiwhy-dating-app/api/presenter"
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID                 uuid.UUID
	UserID             uuid.UUID
	PremiumPackageCode string
	CreatedAt          time.Time
}

func BuildCreateRequest(request presenter.PurchasePremiumPackageRequest) Order {
	return Order{
		ID:                 uuid.New(),
		UserID:             request.UserID,
		PremiumPackageCode: request.PremiumPackageCode,
		CreatedAt:          time.Now(),
	}
}

const (
	PremiumPackageCodeUnlimitedSwipe  = "PP-UNSW"
	PremiumPackageCodeVerifiedAccount = "PP-VEAC"
)
