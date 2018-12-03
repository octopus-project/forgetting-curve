package handler 

import (
	"encoding/json"
	"net/http"

	"go.octopus-project.net/forgetting-curve/internal/cards"
)


func ListCardsHandler(w http.ResponseWriter, r *http.Request) {
	output, err := json.Marshal(cards.All())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	SendResponse(w, output)
}