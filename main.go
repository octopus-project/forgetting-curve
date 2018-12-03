package main


import (
	"log"
	"time"
	"encoding/json"
	"net/http"

	"go.octopus-project.net/forgetting-curve/cmd/api/handler"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/cards", handler.InsertCardsHandler).Methods("POST")
	r.HandleFunc("/cards", handler.ListCardsHandler).Methods("GET")
	r.HandleFunc("/cards/{Id}", handler.GetCardHandler).Methods("GET")

	http.Handle("/", r)
	
	log.Fatal( http.ListenAndServe(":8080", nil) )
}



