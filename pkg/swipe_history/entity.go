package swipehistory

import (
	"rizkiwhy-dating-app/api/presenter"
	"time"

	"github.com/google/uuid"
)

type SwipeHistory struct {
	ID        uuid.UUID
	Sender    uuid.UUID
	Receiver  uuid.UUID
	Swipe     string
	CreatedAt time.Time
}

func BuildCreateRequest(request presenter.ImpressPartnerProfileRequest) SwipeHistory {
	return SwipeHistory{
		ID:        uuid.New(),
		Sender:    request.Sender,
		Receiver:  request.Receiver,
		Swipe:     request.Swipe,
		CreatedAt: time.Now(),
	}
}
