package mocks_test

import (
	"errors"

	"github.com/djilousp/toggl_deck/src/domain"
	"github.com/google/uuid"
)

var decks []*domain.Deck = makeDecks()

func makeDecks() []*domain.Deck {
	fDeck1, _ := uuid.Parse("a251071b-662f-44b6-ba11-e24863039c59")
	fDeck2, _ := uuid.Parse("b301071b-662f-44b6-ba11-e24863039c60")
	return []*domain.Deck{
		{
			ID:       fDeck1,
			Shuffled: false,
		},
		{
			ID:       fDeck2,
			Shuffled: true,
		},
	}

}

type InMemoryDeck struct {
	Decks []*domain.Deck
}

func (d *InMemoryDeck) GetById(deckId string) (*domain.Deck, error) {

	for _, deck := range decks {
		dId, _ := uuid.Parse(deckId)
		if deck.ID == dId {
			return deck, nil
		}
	}
	return nil, errors.New("deck not found!")

}

// func (r *InMemoryDeck) Add(newDeck *domain.Deck) (*domain.Deck , error){
// 	decks = append(decks, newDeck)
// 	return newDeck, nil
// }

func NewInMemoryDeck() *InMemoryDeck {
	return &InMemoryDeck{
		Decks: decks,
	}
}
