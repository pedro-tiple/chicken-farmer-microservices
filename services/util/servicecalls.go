package util

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

const userServiceURL = "http://192.168.99.100:31479/users"
const chickenServiceURL = "http://192.168.99.100:31479/chickens"

func SpendGoldEggs(_userId primitive.ObjectID, _amount uint) error {
	request, err := BuildRequest(
		"POST",
		fmt.Sprintf("%s%s", userServiceURL, "/spendGoldEggs"),
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

func AddFreeChicken(_barnId primitive.ObjectID, _barnOwnerId primitive.ObjectID) error {
	request, err := BuildRequest(
		"GET",
		fmt.Sprintf("%s/buy/%s", chickenServiceURL, _barnId.Hex()),
		nil,
		_barnOwnerId,
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
		return errors.New("couldn't add free chicken")
	}

	return nil
}
