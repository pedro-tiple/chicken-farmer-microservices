package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	usersvc "ptiple/user-svc"
	"ptiple/util"
)

// This isn't proper authentication and the JWT token has no secure properties other than "we know it was generated by our server".
// But for the purposes of this project, it is enough to showcase securing API endpoints with an auth middleware
func (api Api) loginHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := primitive.ObjectIDFromHex(vars["userId"])
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	_, err = api.DB.GetUser(userId)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// login with new user, create it
		_, err := usersvc.New(userId, api.DB)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	tokenString, err := util.GenerateJwtToken(userId, false)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, tokenString)
}

func (api Api) getGoldEggsHandler(w http.ResponseWriter, r *http.Request) {
	jwtToken := r.Context().Value("jwtToken").(*util.JwtToken)

	user, err := api.DB.GetUser(jwtToken.UserId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	_, _ = fmt.Fprintf(w, `{"goldEggCount": %d}`, user.GoldEggs)
}

func (api Api) spendGoldEggsHandler(w http.ResponseWriter, r *http.Request) {
	userId, amount, err := api.initializeApiRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	// TODO use some kind of "select for update" otherwise we can have race conditions and double spending
	user, err := api.DB.GetUser(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	user.DB = api.DB

	err = user.RemoveGoldEggs(amount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (api Api) wsHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		// allow origin to be different from server TODO check if necessary with kubernetes
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer connection.Close()

	util.ListenToTimeUpdates(api.Redis, func(day uint) error {
		return connection.WriteJSON(day)
	})
}

func (api Api) initializeApiRequest(r *http.Request) (primitive.ObjectID, uint, error) {
	var userId primitive.ObjectID

	jwtToken := r.Context().Value("jwtToken").(*util.JwtToken)
	if !jwtToken.IsService {
		return userId, 0, errors.New("request must come from another service")
	}

	requestBody := &struct {
		UserId string `json:"userId"`
		Amount uint   `json:"amount"`
	}{}
	err := json.NewDecoder(r.Body).Decode(requestBody)
	if err != nil {
		return userId, 0, err
	}

	userId, err = primitive.ObjectIDFromHex(requestBody.UserId)
	if err != nil {
		return userId, 0, err
	}

	return userId, requestBody.Amount, nil
}