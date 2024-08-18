package relationshiptype

import (
	"time"
)

type RelationshipType struct {
	ID        int
	Name      string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
