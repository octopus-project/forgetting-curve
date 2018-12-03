package main


import (
	"log"
	"net/http"

	"go.octopus-project.net/forgetting-curve/cmd/api/handler"
	"github.com/gorilla/mux"
)


func main() {

	log.Println("[START] Start Application.")
	
	r := mux.NewRouter()
	r.HandleFunc(handler.InsertCardsRoute, 	handler.InsertCardsHandler).Methods("POST")
	r.HandleFunc(handler.ListCardsRoute, 	handler.ListCardsHandler).Methods("GET")
	r.HandleFunc(handler.GetCardRoute, 		handler.GetCardHandler).Methods("GET")
	r.HandleFunc(handler.CompleteCardRoute, handler.CompleteCardHandler).Methods("GET")
	r.HandleFunc(handler.FailedCardRoute, 	handler.FailedCardHandler).Methods("GET")
	
	http.Handle("/", r)
	
	log.Println("[START] Application is now listening on port 8080.")
	log.Fatal( http.ListenAndServe(":8080", nil) )
}



