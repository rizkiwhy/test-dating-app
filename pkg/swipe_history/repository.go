package swipehistory

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SwipeHistoryRepositoryImpl interface {
	GetTodayBySender(id uuid.UUID) (swipeHistories []SwipeHistory, err error)
	Create(swipeHistory SwipeHistory) (err error)
}

type SwipeHistoryRepository struct {
	DB *gorm.DB
}

func NewRepository(DB *gorm.DB) SwipeHistoryRepositoryImpl {
	return &SwipeHistoryRepository{DB: DB}
}

func (r *SwipeHistoryRepository) GetTodayBySender(id uuid.UUID) (swipeHistories []SwipeHistory, err error) {
	now := time.Now()
	year, month, day := now.Date()
	startAt := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	endAt := time.Date(year, month, day, 23, 59, 59, 59, now.Location())

	err = r.DB.Where("sender = ? AND created_at BETWEEN ? AND ?", id, startAt, endAt).Find(&swipeHistories).Error
	if err != nil {
		fmt.Println("err: ", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fiber.NewError(fiber.StatusNotFound, "premium package not found")
		} else {
			err = fiber.NewError(fiber.StatusInternalServerError, ("internal server error"))
		}
	}

	return
}

func (r *SwipeHistoryRepository) Create(swipeHistory SwipeHistory) (err error) {
	err = r.DB.Create(&swipeHistory).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		switch mysqlErr.Number {
		case 1062: // MySQL code for duplicate entry
			// Handle duplicate entry
		case 1452: // MySQL code for foreign key violation
		// Add cases for other specific error codes
		case 1364: // MySQL code for required field is missing
		// Add cases for other specific error codes
		default:
			// Handle other errors
			err = errors.New("internal server error")
		}
	}

	return
}
