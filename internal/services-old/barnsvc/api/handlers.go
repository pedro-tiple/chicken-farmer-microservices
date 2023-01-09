package api

import (
	"encoding/json"
	"errors"
	"net/http"
	barnsvc "ptiple/barnsvc"
	"ptiple/barnsvc/mongodatabase"
	"ptiple/util"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const feedCost = 1
const barnCost = 10
const autoFeederCost = 100
const FeedPerPurchase = 10

func (api Api) getBarnsHandler(w http.ResponseWriter, r *http.Request) {
	jwtToken := r.Context().Value("jwtToken").(*util.JwtToken)
	barns, err := api.DB.GetBarnsOfFarmer(jwtToken.FarmerId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(barns)
}

func (api Api) getBarnHandler(w http.ResponseWriter, r *http.Request) {
	barn, errCode := getBarnFromRequest(api.DB, r)
	if errCode != -1 {
		w.WriteHeader(errCode)
		return
	}

	_ = json.NewEncoder(w).Encode(barn)
}

func (api Api) buyBarnHandler(w http.ResponseWriter, r *http.Request) {
	jwtToken := r.Context().Value("jwtToken").(*util.JwtToken)

	barns, err := api.DB.GetBarnsOfFarmer(jwtToken.FarmerId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if len(barns) > 0 {
		err = api.ServiceCalls.SpendGoldEggs(jwtToken.FarmerId, barnCost)
		if err != nil {
			w.WriteHeader(http.StatusPreconditionFailed)
			return
		}
	}

	barn, err := barnsvc.New(jwtToken.FarmerId, api.DB)
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	err = api.ServiceCalls.AddFreeChicken(barn.Id, barn.BelongsToFarmer)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(barn)
}

func (api Api) buyFeedHandler(w http.ResponseWriter, r *http.Request) {
	jwtToken := r.Context().Value("jwtToken").(*util.JwtToken)

	barn, errCode := getBarnFromRequest(api.DB, r)
	if errCode != -1 {
		w.WriteHeader(errCode)
		return
	}
	barn.DB = api.DB

	err := api.ServiceCalls.SpendGoldEggs(jwtToken.FarmerId, feedCost)
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	err = barn.AddFeed(FeedPerPurchase)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (api Api) spendFeedHandler(w http.ResponseWriter, r *http.Request) {
	barnId, amount, err := initializeApiRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	// TODO use some kind of "select for update" or locking transaction otherwise we can have race conditions and double spending
	barn, err := api.DB.GetBarn(barnId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	barn.DB = api.DB

	err = barn.RemoveFeed(amount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (api Api) buyAutoFeederHandler(w http.ResponseWriter, r *http.Request) {
	jwtToken := r.Context().Value("jwtToken").(*util.JwtToken)

	barn, errCode := getBarnFromRequest(api.DB, r)
	if errCode != -1 {
		w.WriteHeader(errCode)
		return
	}

	// can't buy if already bought
	if barn.AutoFeeder {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	err := api.ServiceCalls.SpendGoldEggs(jwtToken.FarmerId, autoFeederCost)
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	barn.AutoFeeder = true

	err = api.DB.UpdateBarn(*barn)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getBarnFromRequest(_db mongodatabase.IMongoDatabase, _r *http.Request) (*barnsvc.Barn, int) {
	jwtToken := _r.Context().Value("jwtToken").(*util.JwtToken)

	vars := mux.Vars(_r)
	barnId, err := primitive.ObjectIDFromHex(vars["barnId"])
	if err != nil {
		return nil, http.StatusPreconditionFailed
	}

	barn, err := _db.GetBarn(barnId)
	if err != nil {
		return nil, http.StatusNotFound
	}

	if barn.BelongsToFarmer != jwtToken.FarmerId {
		return nil, http.StatusForbidden
	}

	return barn, -1
}

func initializeApiRequest(r *http.Request) (primitive.ObjectID, uint, error) {
	var barnId primitive.ObjectID

	jwtToken := r.Context().Value("jwtToken").(*util.JwtToken)
	if !jwtToken.IsService {
		return barnId, 0, errors.New("request must come from another service")
	}

	requestBody := &struct {
		BarnId string `json:"barnId"`
		Amount uint   `json:"amount"`
	}{}
	err := json.NewDecoder(r.Body).Decode(requestBody)
	if err != nil {
		return barnId, 0, err
	}

	barnId, err = primitive.ObjectIDFromHex(requestBody.BarnId)
	if err != nil {
		return barnId, 0, err
	}

	return barnId, requestBody.Amount, nil
}
