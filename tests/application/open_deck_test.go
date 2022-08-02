package application_test

import (
	"testing"

	"github.com/djilousp/toggl_deck/src/application"
	mocks_test "github.com/djilousp/toggl_deck/tests/mocks"
	"github.com/stretchr/testify/assert"
)


func MakeDeckApplicationSut() (*application.DeckApplication) {
	deckRepo := mocks_test.NewInMemoryDeck()
	sut := application.NewDeckApplication(deckRepo)
	return sut
}

func TestOpenDeck_Should_Fail_For_Invalid_Id(t *testing.T) {
	sut := MakeDeckApplicationSut()
	_, err := sut.OpenDeck("deck_id") 
	assert.Error(t,err)
}

func TestRandom(t *testing.T){
	assert.Equal(t,true,true)
}