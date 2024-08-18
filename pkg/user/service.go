package user

import (
	"errors"
	"fmt"
	"rizkiwhy-dating-app/api/presenter"
	"rizkiwhy-dating-app/config"
	relationshiptype "rizkiwhy-dating-app/pkg/relationship_type"
	swipehistory "rizkiwhy-dating-app/pkg/swipe_history"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type UserServiceImpl interface {
	RegisterUser(request presenter.RegisterUserRequest) (err error)
	LoginUser(request presenter.LoginUserRequest) (t string, err error)
	SearchPartnerProfile(request presenter.SearchPartnerProfileRequest) (response *presenter.SearchPartnerProfileResponse, err error)
	ImpressPartnerProfile(request presenter.ImpressPartnerProfileRequest) (err error)
}

type UserService struct {
	UserRepository             *UserRepository
	RelationshipTypeRepository *relationshiptype.RelationshipTypeRepository
	SwipeHistoryRepository     *swipehistory.SwipeHistoryRepository
}

func NewService(
	userRepository *UserRepository,
	relationshipTypeRepository *relationshiptype.RelationshipTypeRepository,
	swipeHistoryRepository *swipehistory.SwipeHistoryRepository,
) UserServiceImpl {
	return &UserService{
		UserRepository:             userRepository,
		RelationshipTypeRepository: relationshipTypeRepository,
		SwipeHistoryRepository:     swipeHistoryRepository,
	}
}

func (s *UserService) RegisterUser(request presenter.RegisterUserRequest) (err error) {
	totalUser := s.UserRepository.CountByFilter(BuildCountFilter(request))

	if totalUser > 0 {
		return fiber.NewError(fiber.StatusConflict, "username or email already taken")
	}

	hashedPassword, err := request.HashPassword()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	err = s.UserRepository.Create(BuildCreateRequest(request, hashedPassword))

	return
}

func (s *UserService) LoginUser(request presenter.LoginUserRequest) (t string, err error) {
	user, err := s.UserRepository.GetByEmail(request.Email)
	if err != nil {
		return t, err
	}

	err = user.CompareHashAndPassword(request)
	if err != nil {
		return t, errors.New("invalid credentials")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["iat"] = time.Now().Unix()

	t, err = token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return t, errors.New("internal server error")
	}

	return
}

func (s *UserService) SearchPartnerProfile(request presenter.SearchPartnerProfileRequest) (response *presenter.SearchPartnerProfileResponse, err error) {
	user, err := s.UserRepository.GetByID(request.UserID)
	if err != nil {
		return
	}

	partnerFilter := user.BuildPartnerFilter()
	fmt.Printf("partnerFilter:\n%+v\n", partnerFilter)

	swipeHistories, err := s.SwipeHistoryRepository.GetTodayBySender(request.UserID)
	if err != nil {
		return
	}

	totalSwipes := len(swipeHistories)
	if totalSwipes > 9 && !user.IsUnlimitedSwipe {
		return response, fiber.NewError(fiber.StatusUnprocessableEntity, "you have already swiped 10 times today")
	}

	for _, swipeHistory := range swipeHistories {
		partnerFilter.NotInUserIDs = append(partnerFilter.NotInUserIDs, swipeHistory.Receiver)
	}

	partner, err := s.UserRepository.GetPartnerByFilter(partnerFilter)
	if err != nil {
		return
	}
	if partner == nil {
		return response, fiber.NewError(fiber.StatusNotFound, "partner not found")
	}

	response = partner.BuildPartnerProfileResponse()
	relationshipType, err := s.RelationshipTypeRepository.GetByID(partner.RelationshipTypeID)
	if err != nil {
		return
	}
	response.RelationshipType = relationshipType.Name

	return
}

func (s *UserService) ImpressPartnerProfile(request presenter.ImpressPartnerProfileRequest) (err error) {
	user, err := s.UserRepository.GetByID(request.Sender)
	if err != nil {
		return
	}

	swipeHistories, err := s.SwipeHistoryRepository.GetTodayBySender(request.Sender)
	if err != nil {
		return
	}

	totalSwipes := len(swipeHistories)
	if totalSwipes > 10 && !user.IsUnlimitedSwipe {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "you have already swiped 10 times today")
	}

	err = s.SwipeHistoryRepository.Create(swipehistory.BuildCreateRequest(request))
	if err != nil {
		return
	}

	return
}
