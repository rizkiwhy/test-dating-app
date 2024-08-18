package order

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type OrderRepositoryImpl interface {
	Create(order Order) (err error)
}

type OrderRepository struct {
	DB *gorm.DB
}

func NewRepository(DB *gorm.DB) OrderRepositoryImpl {
	return &OrderRepository{DB: DB}
}

func (r *OrderRepository) Create(order Order) (err error) {
	err = r.DB.Create(&order).Error
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
