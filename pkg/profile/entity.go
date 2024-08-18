package profile

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	ID              uuid.UUID
	Name            string
	Birthday        string
	Gender          string
	IsGenderVisible bool
	SrcPhoto        string
	CreatedAt       time.Time
	UpdatedAt       *time.Time
	DeletedAt       *time.Time
}
