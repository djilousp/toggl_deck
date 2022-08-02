package dtos

import (
	"time"

	"github.com/google/uuid"
)
type DeckModelMapping struct {
	ID  		uuid.UUID  `gorm:"primary_key" `
	Shuffled	bool       `gorm:"boolean;default:false"`
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
}
type DeckInputDTO struct {
	ID  		uuid.UUID `json:"deck_id"`
	Shuffled	bool      `json:"shuffled"`
}

type DeckOutPutDTO struct {
	ID  		uuid.UUID `json:"deck_id"`
	Shuffled	bool      `json:"shuffled"`
	Remaining	int64     `json:"remaining"`
}