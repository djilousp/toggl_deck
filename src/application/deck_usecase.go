package application

import (
	dtos "github.com/djilousp/toggl_deck/src/application/DTO"
	"github.com/djilousp/toggl_deck/src/application/protocols"
)

type DeckApplication struct {
	deckRepository protocols.DeckRepository
}



func NewDeckApplication(repo protocols.DeckRepository) *DeckApplication {

	return &DeckApplication{
		deckRepository: repo,
	}
}

func (da *DeckApplication) OpenDeck(deckId string) (*dtos.DeckOutPutDTO, error){
	deck, err := da.deckRepository.GetById(deckId)
	return deck.ToOutPutDTO(), err
}
