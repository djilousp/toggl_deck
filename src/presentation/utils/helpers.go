package utils

import presentation_protocols "github.com/djilousp/toggl_deck/src/presentation/protocols"



func Ok(jsonData []byte) *presentation_protocols.HttpResponse {
	return &presentation_protocols.HttpResponse{
		StatusCode: 200,
		JsonBody:   jsonData,
	}
}

func BadRequest(err error) *presentation_protocols.HttpResponse {
	return &presentation_protocols.HttpResponse{
		StatusCode: 400,
		JsonBody:   []byte(err.Error()),
	}
}

func ServerError(err error) *presentation_protocols.HttpResponse {
	return &presentation_protocols.HttpResponse{
		StatusCode: 500,
		JsonBody:   []byte(err.Error()),
	}
}
