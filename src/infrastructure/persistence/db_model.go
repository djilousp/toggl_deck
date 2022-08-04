package persistence

import (
	"time"

	"github.com/google/uuid"
)

type Deck struct {
	ID        uuid.UUID `gorm:"primary_key" `
	Shuffled  bool      `gorm:"boolean;default:false"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}


type Card struct {
	ID        int64     `gorm:"primary_key" `
	Suit  	  string    `gorm:"type:varchar(8);"`
	Value 	  string	`gorm:"type:varchar(3);"`
	Code 	  string	`gorm:"type:varchar(3);"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type DecksCards struct {
	ID        int64     `gorm:"primary_key" `
	DeckId    uuid.UUID 
	CardId 	  int64		
	Order     int		`gorm:"type:integer"`
	Drawn     bool		`gorm:"boolean;default:false"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Deck      Deck		`gorm:"foreignKey:DeckId;references:DeckId;constraint:OnDelete:CASCADE"`
	Card      Card		`gorm:"foreignKey:CardId;references:CardId;constraint:OnDelete:CASCADE"`
}