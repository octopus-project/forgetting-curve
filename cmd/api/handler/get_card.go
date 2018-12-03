package handler

import (
	"net/http"
	"encoding/json"

	"go.octopus-project.net/forgetting-curve/internal/cards"
	"github.com/satori/go.uuid"
	"github.com/gorilla/mux"
)

func GetCardHandler(w http.ResponseWriter, r *http.Request) {
	vars 	:= mux.Vars(r)
	var_id  := vars["Id"]
	Id, _   := uuid.FromString( var_id ) 


	if cards.Exists(Id) == false {
		SendNotFound(w)
		return
	}

	output, err := json.Marshal( cards.Get(Id) )
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	SendResponse(w, output)
}