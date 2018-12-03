package main


import (
	"log"
	"net/http"

	"go.octopus-project.net/forgetting-curve/cmd/api/handler"
	"github.com/gorilla/mux"
)


func main() {

	r := mux.NewRouter()
	r.HandleFunc(handler.InsertCardsRoute, 	handler.InsertCardsHandler).Methods("POST")
	r.HandleFunc(handler.ListCardsRoute, 	handler.ListCardsHandler).Methods("GET")
	r.HandleFunc(handler.GetCardRoute, handler.GetCardHandler).Methods("GET")
	r.HandleFunc(handler.CompleteCardRoute, handler.CompleteCardHandler).Methods("GET")

	http.Handle("/", r)
	
	log.Fatal( http.ListenAndServe(":8080", nil) )
}



