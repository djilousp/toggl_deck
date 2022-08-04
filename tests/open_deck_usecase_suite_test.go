package application_test

import (
	"errors"
	"testing"

	"github.com/djilousp/toggl_deck/src/application"
	dtos "github.com/djilousp/toggl_deck/src/application/DTO"
	"github.com/djilousp/toggl_deck/src/domain"
	mocks_test "github.com/djilousp/toggl_deck/tests/mocks"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func MakeDeckApplicationSut() *application.DeckApplication {
	deckRepo := mocks_test.NewInMemoryDeck()
	sut := application.NewDeckApplication(deckRepo)
	return sut
}

func MakeFakeDeck() *dtos.DeckOutPutDTO {
	fakeUUID, _ := uuid.Parse("a251071b-662f-44b6-ba11-e24863039c59")
	fakeDeck := *domain.NewDeck(fakeUUID, false)
	return fakeDeck.ToOutPutDTO()
}

func TestTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tests Suite")
}

var _ = Describe("Open Deck Usecase", func() {
	Context("Open Deck should return error for Invalid ID", func() {
		It("returns error", func() {
			sut := MakeDeckApplicationSut()
			_, err := sut.OpenDeck("adsadasd")
			Ω(err).Error().Should(HaveOccurred())
			Ω(err).Should(Equal(errors.New("deck not found!")))
		})
	})
	Context("Open Deck should return a deck for a given  UUID", func() {
		It("returns a valid deck", func() {
			sut := MakeDeckApplicationSut()
			deck, _ := sut.OpenDeck("a251071b-662f-44b6-ba11-e24863039c59")
			fakeDeck := MakeFakeDeck()
			Ω(*deck).To(Equal(*fakeDeck))
		})
	})
})
