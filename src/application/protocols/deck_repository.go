package protocols

import (
	"github.com/djilousp/toggl_deck/src/domain"
)

type DeckRepository interface {
	//GetById(deckId string) (*dtos.DeckOutPutDTO,error)
	GetById(deckId string) (*domain.Deck, error)
	//Add(deck *domain.Deck) (*domain.Deck,error)
}
