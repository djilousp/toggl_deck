package mocks_test

import (
	"errors"

	"github.com/djilousp/toggl_deck/src/domain"
	"github.com/google/uuid"
)
var decks []*domain.Deck

type InMemoryDeck struct{
	Decks []*domain.Deck
}

func (d *InMemoryDeck)GetById(deckId string) (*domain.Deck , error){
	for _, deck := range decks {
		dId, _ := uuid.Parse(deckId)
		if deck.ID == dId  {
			return deck, nil
		}
	}
	return nil, errors.New("deck not found!")
	
}

// func (r *InMemoryDeck) Create(deckId string) (*dtos.DeckOutPutDTO , error){
// 	deck, err:= r.GetById(deckId)
// 	return deck,err
// }
func NewInMemoryDeck() *InMemoryDeck{
	return &InMemoryDeck{
		Decks: decks,
	}
}