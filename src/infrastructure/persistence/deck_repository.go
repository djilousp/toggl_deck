package persistence

import (
	"errors"

	"github.com/djilousp/toggl_deck/src/application/protocols"
	"github.com/djilousp/toggl_deck/src/domain"
	"github.com/jinzhu/gorm"
)

type DeckRepo struct {
	db *gorm.DB
}

func NewDeckRepository(db *gorm.DB) *DeckRepo {
	return &DeckRepo{db}
}

//Deck implements the protocols.DeckRepository interface
var _ protocols.DeckRepository = &DeckRepo{}

func (r *DeckRepo) GetById(deckId string) (*domain.Deck, error) {
	var deck domain.Deck
	err := r.db.Debug().Where("id = ?", deckId).Take(&deck).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("deck not found")
	}
	return &deck, nil
}
