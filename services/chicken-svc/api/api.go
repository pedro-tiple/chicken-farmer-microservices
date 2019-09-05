package api

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"ptiple/chicken-svc/mongodatabase"
)

type Api struct {
	DB *mongodatabase.MongoDatabase
}

func Start() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mongodb, err := mongodatabase.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	var api = Api{&mongodb}

	api.setupHandlers()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (api Api) setupHandlers() {
	r := mux.NewRouter()
	r.HandleFunc("/chickens/barn/{barnId}", api.getChickensOfBarnHandler).Methods("GET")
	r.HandleFunc("/chickens/buy/{barnId}", api.buyChickenHandler).Methods("GET")
	r.HandleFunc("/chickens/{chickenId}", api.getChickenHandler).Methods("GET")
	r.HandleFunc("/chickens/{chickenId}/feed", api.feedChickenHandler).Methods("GET")
	r.HandleFunc("/chickens/{chickenId}/sell", api.sellChickenHandler).Methods("GET")

	http.Handle("/", r)
}
