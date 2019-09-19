package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"math/rand"
	"net/http"
	chickensvc "ptiple/chicken-svc"
	"ptiple/util"
)

const barnServiceURL = "http://192.168.99.100:31479/barns"
const userServiceURL = "http://192.168.99.100:31479/users"
const chickenCostInGoldEggs = 1
const chickenFeedConsumption = 1

func (api Api) getChickensOfUserHandler(w http.ResponseWriter, r *http.Request) {
	jwtToken := r.Context().Value("jwtToken").(*util.JwtToken)
	barn, err := api.DB.GetChickensOfUser(jwtToken.UserId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	_ = json.NewEncoder(w).Encode(barn)
}

func (api Api) feedChickenHandler(w http.ResponseWriter, r *http.Request) {
	jwtToken := r.Context().Value("jwtToken").(*util.JwtToken)

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

	if chicken.BelongsToUser != jwtToken.UserId {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	laidGoldEgg, err := api.feedChicken(chicken)
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	_, err = w.Write([]byte(fmt.Sprintf(`{"laidGoldEgg": %t, "restingUntil": %d}`, laidGoldEgg, chicken.RestingUntil)))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

type BulkFeedChickenResult struct {
	Id           string `json:"id"`
	LaidGoldEgg  bool   `json:"laidGoldEgg"`
	RestingUntil uint   `json:"restingUntil"`
}

func (api Api) bulkFeedChickenHandler(w http.ResponseWriter, r *http.Request) {
	jwtToken := r.Context().Value("jwtToken").(*util.JwtToken)

	requestBody := &struct {
		ChickenIds []string `json:"chickenIds"`
	}{}
	err := json.NewDecoder(r.Body).Decode(requestBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var results = []BulkFeedChickenResult{}
	// TODO instead of handling each chicken individually, also do it in bulk
	for _, chickenIdStr := range requestBody.ChickenIds {
		chickenId, err := primitive.ObjectIDFromHex(chickenIdStr)
		if err != nil {
			continue
		}

		chicken, err := api.DB.GetChicken(chickenId)
		if err != nil {
			continue
		}

		if chicken.BelongsToUser != jwtToken.UserId {
			continue
		}

		laidGoldEgg, err := api.feedChicken(chicken)
		if err != nil {
			continue
		}

		results = append(
			results,
			BulkFeedChickenResult{
				Id:           chicken.Id.Hex(),
				LaidGoldEgg:  laidGoldEgg,
				RestingUntil: chicken.RestingUntil,
			},
		)
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(results)
}

func (api Api) feedChicken(chicken *chickensvc.Chicken) (bool, error) {

	currentDay, err := api.DB.GetDay()
	if err != nil {
		return false, err
	}

	if currentDay <= chicken.RestingUntil {
		return false, errors.New("chicken still resting")
	}

	if err = spendFeed(chicken.BelongsToUser, chicken.BelongsToBarn, chickenFeedConsumption); err != nil {
		return false, err
	}

	chicken.DB = api.DB
	chicken.Rng = rand.Intn
	laidGoldEgg, err := chicken.Feed(currentDay)
	if err != nil {
		return false, err
	}

	if laidGoldEgg {
		err := api.Redis.Publish(
			"chicken-updates",
			fmt.Sprintf(`{"userId": "%s","chickenId": "%s","event": "laidGoldEgg"}`, chicken.BelongsToUser.Hex(), chicken.Id.Hex()),
		).Err()
		if err != nil {
			log.Println("failed publishing to redis", err)
		}
	}

	return laidGoldEgg, nil
}

func (api Api) buyChickenHandler(w http.ResponseWriter, r *http.Request) {
	jwtToken := r.Context().Value("jwtToken").(*util.JwtToken)

	vars := mux.Vars(r)
	barnId, err := primitive.ObjectIDFromHex(vars["barnId"])
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	if !checkBarnOwnerShip(jwtToken.UserId, barnId.Hex()) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	currentDay, err := api.DB.GetDay()
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	if !jwtToken.IsService {
		err = spendGoldEggs(jwtToken.UserId, chickenCostInGoldEggs)
		if err != nil {
			w.WriteHeader(http.StatusPreconditionFailed)
			return
		}
	}

	chicken, err := chickensvc.New(barnId, jwtToken.UserId, currentDay, api.DB, rand.Intn)
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(chicken)
}

func (api Api) sellChickenHandler(w http.ResponseWriter, r *http.Request) {
	jwtToken := r.Context().Value("jwtToken").(*util.JwtToken)

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

	if chicken.BelongsToUser != jwtToken.UserId {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if err := api.DB.RemoveChicken(chickenId); err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func spendFeed(_userId primitive.ObjectID, _barnId primitive.ObjectID, _amount uint) error {
	request, err := util.BuildRequest(
		"POST",
		fmt.Sprintf("%s/spendFeed", barnServiceURL),
		struct {
			BarnId string `json:"barnId"`
			Amount uint   `json:"amount"`
		}{
			BarnId: _barnId.Hex(),
			Amount: _amount,
		},
		_userId,
	)
	if err != nil {
		return err
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return errors.New("couldn't spend the requested amount")
	}

	return nil
}

func spendGoldEggs(_userId primitive.ObjectID, _amount uint) error {
	request, err := util.BuildRequest(
		"POST",
		fmt.Sprintf("%s/spendGoldEggs", userServiceURL),
		struct {
			UserId string `json:"userId"`
			Amount uint   `json:"amount"`
		}{
			UserId: _userId.Hex(),
			Amount: _amount,
		},
		_userId,
	)
	if err != nil {
		return err
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return errors.New("couldn't spend the requested amount")
	}

	return nil
}

func checkBarnOwnerShip(_owner primitive.ObjectID, _barnId string) bool {
	request, err := util.BuildRequest(
		"GET",
		fmt.Sprintf("%s/%s", barnServiceURL, _barnId),
		nil,
		_owner,
	)
	if err != nil {
		return false
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil || response.StatusCode != http.StatusOK {
		return false
	}

	requestBody := &struct {
		BelongsToUser string `json:"belongsToUser"`
	}{}
	err = json.NewDecoder(response.Body).Decode(requestBody)
	if err != nil {
		return false
	}

	if requestBody.BelongsToUser == _owner.Hex() {
		return true
	}

	return false
}
