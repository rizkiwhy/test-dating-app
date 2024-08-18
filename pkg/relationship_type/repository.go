package relationshiptype

import (
	"errors"

	"gorm.io/gorm"
)

type RelationshipTypeRepositoryImpl interface {
	GetByID(id int) (relationshipType *RelationshipType, err error)
}

type RelationshipTypeRepository struct {
	DB *gorm.DB
}

func NewRepository(DB *gorm.DB) RelationshipTypeRepositoryImpl {
	return &RelationshipTypeRepository{DB: DB}
}

func (r *RelationshipTypeRepository) GetByID(id int) (relationshipType *RelationshipType, err error) {
	err = r.DB.Where("id = ?", id).First(&relationshipType).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("user not found")
		} else {
			err = errors.New("internal server error")
		}
	}

	return
}
