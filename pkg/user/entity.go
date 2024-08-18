package user

import (
	"fmt"
	"rizkiwhy-dating-app/api/presenter"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                 uuid.UUID
	Username           string
	Email              string
	Password           string
	IsActive           bool
	CreatedAt          time.Time
	UpdatedAt          *time.Time
	DeletedAt          *time.Time
	Name               string
	Birthday           time.Time
	Gender             string
	IsGenderVisible    bool
	IsUnlimitedSwipe   bool
	IsVerifiedAccount  bool
	RelationshipTypeID int
	Interested         string
}

// CalculateAge calculates the age based on the birthday
func (user User) CalculateAge() int {
	// Get the current date
	now := time.Now()
	// Calculate age based on the birthday
	age := now.Year() - user.Birthday.Year()

	// Adjust age if the birthday hasn't occurred yet this year
	if now.YearDay() < user.Birthday.YearDay() {
		age--
	}

	return age
}

func (user User) BuildPartnerProfileResponse() *presenter.SearchPartnerProfileResponse {
	response := &presenter.SearchPartnerProfileResponse{
		PartnerID:          user.ID,
		IsVerifiedAccount:  user.IsVerifiedAccount,
		Interested:         user.Interested,
		RelationshipTypeID: user.RelationshipTypeID,
	}

	if user.IsGenderVisible {
		response.Gender = &user.Gender
	}

	response.Age = fmt.Sprintf("%d years old", user.CalculateAge())

	return response
}

func (user User) BuildPartnerFilter() (filter GetPartnerFilter) {
	filter.NotInUserIDs = append(filter.NotInUserIDs, user.ID)

	switch {
	case user.Gender == "man" && user.Interested == "women":
		filter.Genders = append(filter.Genders, "woman")
		filter.InterestedInUsers = append(filter.InterestedInUsers, "men")
	case user.Gender == "woman" && user.Interested == "men":
		filter.Genders = append(filter.Genders, "man")
		filter.InterestedInUsers = append(filter.InterestedInUsers, "women")
	default:
		filter.Genders = []string{"man", "woman"}
		filter.InterestedInUsers = []string{"everyone", "men", "women"}
	}

	switch user.RelationshipTypeID {
	case 1, 2:
		filter.RelationshipTypeIDs = []int{1, 2, 3}
	case 3:
		filter.RelationshipTypeIDs = []int{2, 3, 4}
	case 4:
		filter.RelationshipTypeIDs = []int{3, 4, 5}
	case 5:
		filter.RelationshipTypeIDs = []int{4, 5, 6}
	case 6:
		filter.RelationshipTypeIDs = []int{1, 2, 3, 4, 5, 6}
	}

	filter.BirthdaysRange = []time.Time{user.Birthday.AddDate(-5, 0, 0), user.Birthday.AddDate(5, 0, 0)}

	return
}

type CountFilter struct {
	Username *string
	Email    *string
}

type GetPartnerFilter struct {
	NotInUserIDs        []uuid.UUID
	Genders             []string
	BirthdaysRange      []time.Time
	RelationshipTypeIDs []int
	InterestedInUsers   []string
}

func BuildCountFilter(request presenter.RegisterUserRequest) CountFilter {
	return CountFilter{
		Username: &request.Username,
		Email:    &request.Email,
	}
}

func BuildCreateRequest(request presenter.RegisterUserRequest, hashedPassword string) User {
	parsedBirthday, _ := time.Parse("2006-01-02", request.Birthday)
	return User{
		ID:              uuid.New(),
		Username:        request.Username,
		Email:           request.Email,
		Password:        hashedPassword,
		IsActive:        true,
		CreatedAt:       time.Now(),
		Name:            request.Name,
		Birthday:        parsedBirthday,
		Gender:          request.Gender,
		IsGenderVisible: request.IsGenderVisible,
	}
}

func (user User) CompareHashAndPassword(request presenter.LoginUserRequest) (err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return err
	}

	return
}
