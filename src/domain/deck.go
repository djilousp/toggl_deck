package domain

import (
	dtos "github.com/djilousp/toggl_deck/src/application/DTO"
	"github.com/google/uuid"
)

type Deck struct {
	ID 			uuid.UUID
	Shuffled 	bool
	Cards 		[]*Cards	
}

func (d *Deck) ToOutPutDTO() (*dtos.DeckOutPutDTO) {
		return &dtos.DeckOutPutDTO{
			ID: d.ID,
			Shuffled: d.Shuffled,
			//Remaining: len(d.Cards),
		}
}

func (d *Deck) GetRemaining() {

}
func NewDeck(id uuid.UUID, shuffled bool) *Deck{
	return &Deck{
		ID: id,
		Shuffled: shuffled,
	}
}

