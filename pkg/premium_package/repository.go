package premiumpackage

import (
	"errors"

	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type PremiumPackageRepositoryImpl interface {
	GetAll() (premiumPackage []PremiumPackage, err error)
}

type PremiumPackageRepository struct {
	DB *gorm.DB
}

func NewRepository(DB *gorm.DB) PremiumPackageRepositoryImpl {
	return &PremiumPackageRepository{DB: DB}
}

func (r *PremiumPackageRepository) GetAll() (premiumPackages []PremiumPackage, err error) {
	err = r.DB.Where(&PremiumPackage{DeletedAt: nil, IsActive: true}).Find(&premiumPackages).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fiber.NewError(fiber.StatusNotFound, "premium package not found")
		} else {
			err = fiber.NewError(fiber.StatusInternalServerError, ("internal server error"))
		}
	}

	return
}
