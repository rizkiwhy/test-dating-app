package premiumpackage

import "time"

type PremiumPackage struct {
	Code      string
	Name      string
	Price     float64
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
