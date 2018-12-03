package handler 

import (
	"net/http"
	"encoding/json"

	"github.com/satori/go.uuid"
	"github.com/gorilla/mux"
	"go.octopus-project.net/forgetting-curve/internal/cards"
)

const CompleteCardRoute = "/cards/{Id}/complete"

func CompleteCardHandler(w http.ResponseWriter, r *http.Request) {

	vars 	:= mux.Vars(r)
	var_id  := vars["Id"]
	Id, _   := uuid.FromString( var_id ) 


	if cards.Exists(Id) == false {
		SendNotFound(w)
		return
	}

	Entry := cards.Get(Id)
	Entry = cards.Step(Entry, Entry.Step+1)

	output, err := json.Marshal( Entry )
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	SendResponse(w, output)
}