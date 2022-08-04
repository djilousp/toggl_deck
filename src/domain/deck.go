package domain

import (
	dtos "github.com/djilousp/toggl_deck/src/application/DTO"
	"github.com/google/uuid"
)

type Deck struct {
	ID         	uuid.UUID
	Shuffled	bool
	Cards		[]*Card
	DrawnCards	int
}

func (d *Deck) ToOutPutDTO() *dtos.DeckOutPutDTO {
	rm := d.GetRemaining()
	return &dtos.DeckOutPutDTO{
		ID:        d.ID,
		Shuffled:  d.Shuffled,
		Remaining: rm,
	}
}


func (d *Deck) GetRemaining() int {
	return len(d.Cards) - d.DrawnCards
}
func NewDeck(id uuid.UUID, shuffled bool) *Deck {
	return &Deck{
		ID:       id,
		Shuffled: shuffled,
	}
}
