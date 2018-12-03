package handler

import (
	"net/http"
	"encoding/json"

	"go.octopus-project.net/forgetting-curve/internal/cards"
)


const InsertCardsRoute = "/cards"

func InsertCardsHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var FormatedBody cards.Card

	err := decoder.Decode(&FormatedBody)
	if err != nil {
		panic(err)
	}

	CardToBeSaved := cards.Create(FormatedBody)

	output, err := json.Marshal(CardToBeSaved)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    SendResponse(w, output)
}