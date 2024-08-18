package order

import (
	"rizkiwhy-dating-app/api/presenter"
	"rizkiwhy-dating-app/pkg/user"

	"github.com/gofiber/fiber/v2"
)

type OrderServiceImpl interface {
	PurchasePremiumPackage(request presenter.PurchasePremiumPackageRequest) (err error)
}

type OrderService struct {
	OrderRepository *OrderRepository
	UserRepository  *user.UserRepository
}

func NewService(
	orderRepository *OrderRepository,
	userRepository *user.UserRepository,
) OrderServiceImpl {
	return &OrderService{
		OrderRepository: orderRepository,
		UserRepository:  userRepository,
	}
}

func (o *OrderService) PurchasePremiumPackage(request presenter.PurchasePremiumPackageRequest) (err error) {
	user, err := o.UserRepository.GetByID(request.UserID)
	if err != nil {
		return
	}

	switch request.PremiumPackageCode {
	case PremiumPackageCodeUnlimitedSwipe:
		if user.IsUnlimitedSwipe {
			return fiber.NewError(fiber.StatusConflict, ("premium package already purchased"))
		}
		user.IsUnlimitedSwipe = true

	case PremiumPackageCodeVerifiedAccount:
		if user.IsVerifiedAccount {
			return fiber.NewError(fiber.StatusConflict, ("premium package already purchased"))
		}
		user.IsVerifiedAccount = true
	default:
		return fiber.NewError(fiber.StatusBadRequest, ("invalid premium package code"))
	}

	err = o.OrderRepository.Create(BuildCreateRequest(request))
	if err != nil {
		return
	}

	switch {
	case user.IsUnlimitedSwipe:
		err = o.UserRepository.UpdateUnlimitedSwipe(*user)
	case user.IsVerifiedAccount:
		err = o.UserRepository.UpdateUnlimitedSwipe(*user)
	}

	return
}
