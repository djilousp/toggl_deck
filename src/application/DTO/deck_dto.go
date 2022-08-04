package dtos

import (
	"github.com/google/uuid"
)

type DeckInputDTO struct {
	ID       uuid.UUID `json:"deck_id"`
	Shuffled bool      `json:"shuffled"`
}

type DeckOutPutDTO struct {
	ID        uuid.UUID 	`json:"deck_id"`
	Shuffled  bool      	`json:"shuffled"`
	Remaining int     		`json:"remaining"`
	Cards 	  CardOutPutDTO	`json:"cards"`
}