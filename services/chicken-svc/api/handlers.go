package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"math/rand"
	"net/http"
	chickensvc "ptiple/chicken-svc"
)

func (api Api) getChickenHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chickenId, err := primitive.ObjectIDFromHex(vars["chickenId"])
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	chicken, err := api.DB.GetChicken(chickenId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	_ = json.NewEncoder(w).Encode(chicken)
}

func (api Api) getChickensOfBarnHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	barnId, err := primitive.ObjectIDFromHex(vars["barnId"])
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	chicken, err := api.DB.GetChickensOfBarn(barnId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(chicken)
}

func (api Api) feedChickenHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chickenId, err := primitive.ObjectIDFromHex(vars["chickenId"])
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	chicken, err := api.DB.GetChicken(chickenId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	chicken.DB = api.DB
	chicken.Rng = rand.Intn
	err = chicken.Feed()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (api Api) buyChickenHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	barnId, err := primitive.ObjectIDFromHex(vars["barnId"])
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	// TODO get the current day
	chicken, err := chickensvc.New(1, barnId, api.DB, rand.Intn)
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(chicken)
}

func (api Api) sellChickenHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chickenId, err := primitive.ObjectIDFromHex(vars["chickenId"])
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	// make sure it exists
	_, err = api.DB.GetChicken(chickenId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := api.DB.RemoveChicken(chickenId); err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	w.WriteHeader(http.StatusOK)
}
