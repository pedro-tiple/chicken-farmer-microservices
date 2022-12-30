package util

import (
	"errors"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const farmerServiceURL = "http://192.168.99.100:31479/farmers"
const chickenServiceURL = "http://192.168.99.100:31479/chickens"

func SpendGoldEggs(_farmerId primitive.ObjectID, _amount uint) error {
	request, err := BuildRequest(
		"POST",
		fmt.Sprintf("%s%s", farmerServiceURL, "/spendGoldEggs"),
		struct {
			FarmerId string `json:"farmerId"`
			Amount   uint   `json:"amount"`
		}{
			FarmerId: _farmerId.Hex(),
			Amount:   _amount,
		},
		_farmerId,
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
		return errors.New("couldn't add free farm-old")
	}

	return nil
}
