package controllers

import (
	"encoding/json"
	"errors"

	"github.com/djilousp/toggl_deck/src/application"
	presentation_protocols "github.com/djilousp/toggl_deck/src/presentation/protocols"
	"github.com/djilousp/toggl_deck/src/presentation/utils"
)

type DeckController struct {
	deckApplication *application.DeckApplication
}

func (controller *DeckController) OpenDeck(request *presentation_protocols.HttpRequest) *presentation_protocols.HttpResponse {

	if len(request.Params) == 0 {
		err := errors.New("Blank DeckId!")
		return utils.BadRequest(err)
	}
	deckId := request.Params[0]

	foundDeck, err := controller.deckApplication.OpenDeck(string(deckId))
	if err != nil {
		return utils.ServerError(err)
	}

	jsonOutputDto, err := json.Marshal(foundDeck)
	if err != nil {
		return utils.ServerError(err)
	}

	return utils.Ok(jsonOutputDto)
}

func NewAddUserController(deckApplication *application.DeckApplication) *DeckController {
	return &DeckController{
		deckApplication: deckApplication,
	}
}
