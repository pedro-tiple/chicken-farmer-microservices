package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	barnsvc "ptiple/barn-svc"
	"ptiple/barn-svc/mocks"
	"ptiple/util"
	"strings"
	"testing"
)

func TestHandlers_GetBarnsHandlerEmpty(t *testing.T) {
	userId := primitive.NewObjectID()
	recorder, request := setupRecorderAndRequest(
		httptest.NewRequest("GET", "/barns", nil),
		userId,
		false,
	)
	mongodb, api := setupMongoMockAndAPI(t)

	mongodb.
		EXPECT().
		GetBarnsOfUser(gomock.Eq(userId)).
		Return([]*barnsvc.Barn{}, nil)

	api.getBarnsHandler(recorder, request)
	result := recorder.Result()

	statusCode := result.StatusCode
	if recorder.Result().StatusCode != 200 {
		t.Errorf("Expected StatusCode 200 got %d", statusCode)
	}

	body, _ := ioutil.ReadAll(result.Body)
	trimmedBody := strings.TrimSpace(string(body))
	if trimmedBody != "[]" {
		t.Errorf("Expected [] got '%s'", trimmedBody)
	}
}

func TestHandlers_GetBarnsHandlerWithBarns(t *testing.T) {
	userId := primitive.NewObjectID()
	recorder, request := setupRecorderAndRequest(
		httptest.NewRequest("GET", "/barns", nil),
		userId,
		false,
	)
	mongodb, api := setupMongoMockAndAPI(t)

	barns := []*barnsvc.Barn{
		{
			Id:            primitive.NewObjectID(),
			BelongsToUser: userId,
			Feed:          10,
			AutoFeeder:    false,
			DB:            nil,
		},
		{
			Id:            primitive.NewObjectID(),
			BelongsToUser: userId,
			Feed:          100,
			AutoFeeder:    true,
			DB:            nil,
		},
	}
	mongodb.
		EXPECT().
		GetBarnsOfUser(gomock.Eq(userId)).
		Return(barns, nil)

	api.getBarnsHandler(recorder, request)
	result := recorder.Result()

	statusCode := result.StatusCode
	if recorder.Result().StatusCode != 200 {
		t.Errorf("Expected StatusCode 200 got %d", statusCode)
	}

	barnsResult := []barnsvc.Barn{}
	body, _ := ioutil.ReadAll(result.Body)
	_ = json.Unmarshal(body, &barnsResult)
	if len(barnsResult) != 2 {
		t.Errorf("Expected two barns got %d", len(barnsResult))
	}
}

func TestHandlers_GetBarnHandler(t *testing.T) {
	userId := primitive.NewObjectID()
	barnId := primitive.NewObjectID()

	recorder, request := setupRecorderAndRequest(
		httptest.NewRequest("GET", fmt.Sprintf("/barns/%s", barnId.Hex()), nil),
		userId,
		false,
	)
	request = mux.SetURLVars(request, map[string]string{"barnId": barnId.Hex()})
	mongodb, api := setupMongoMockAndAPI(t)

	mongodb.
		EXPECT().
		GetBarn(gomock.Eq(barnId)).
		Return(&barnsvc.Barn{
			Id:            barnId,
			BelongsToUser: userId,
			Feed:          0,
			AutoFeeder:    false,
			DB:            nil,
		}, nil)

	api.getBarnHandler(recorder, request)
	result := recorder.Result()

	statusCode := result.StatusCode
	if recorder.Result().StatusCode != 200 {
		t.Errorf("Expected StatusCode 200 got %d", statusCode)
	}

	body, _ := ioutil.ReadAll(result.Body)
	trimmedBody := strings.TrimSpace(string(body))
	expected := fmt.Sprintf(`{"id":"%s","belongsToUser":"%s","feed":0,"autoFeeder":false}`, barnId.Hex(), userId.Hex())
	if trimmedBody != expected {
		t.Errorf("Expected %s got '%s'", expected, trimmedBody)
	}
}

func TestHandlers_BuyFreeBarnHandler(t *testing.T) {
	userId := primitive.NewObjectID()

	recorder, request := setupRecorderAndRequest(
		httptest.NewRequest("GET", "/barns/buy", nil),
		userId,
		false,
	)
	mongodb, api := setupMongoMockAndAPI(t)

	api.ServiceCalls.AddFreeChicken = func(_barnId primitive.ObjectID, _barnOwnerId primitive.ObjectID) error {
		return nil
	}

	mongodb.
		EXPECT().
		GetBarnsOfUser(gomock.Eq(userId)).
		Return([]*barnsvc.Barn{}, nil)

	mongodb.
		EXPECT().
		InsertBarn(gomock.Any()).
		Return(&barnsvc.Barn{
			Id:            primitive.NewObjectID(),
			BelongsToUser: userId,
			Feed:          0,
			AutoFeeder:    false,
			DB:            nil,
		}, nil)

	api.buyBarnHandler(recorder, request)
	result := recorder.Result()

	statusCode := result.StatusCode
	if recorder.Result().StatusCode != 200 {
		t.Errorf("Expected StatusCode 200 got %d", statusCode)
	}
}

func TestHandlers_BuyPayedBarnHandler(t *testing.T) {
	userId := primitive.NewObjectID()

	recorder, request := setupRecorderAndRequest(
		httptest.NewRequest("GET", "/barns/buy", nil),
		userId,
		false,
	)
	mongodb, api := setupMongoMockAndAPI(t)

	api.ServiceCalls.SpendGoldEggs = func(_userId primitive.ObjectID, _amount uint) error {
		return nil
	}
	api.ServiceCalls.AddFreeChicken = func(_barnId primitive.ObjectID, _barnOwnerId primitive.ObjectID) error {
		return nil
	}

	mongodb.
		EXPECT().
		GetBarnsOfUser(gomock.Eq(userId)).
		Return([]*barnsvc.Barn{{}, {}}, nil)

	mongodb.
		EXPECT().
		InsertBarn(gomock.Any()).
		Return(&barnsvc.Barn{
			Id:            primitive.NewObjectID(),
			BelongsToUser: userId,
			Feed:          0,
			AutoFeeder:    false,
			DB:            nil,
		}, nil)

	api.buyBarnHandler(recorder, request)
	result := recorder.Result()

	statusCode := result.StatusCode
	if recorder.Result().StatusCode != 200 {
		t.Errorf("Expected StatusCode 200 got %d", statusCode)
	}
}

func TestHandlers_BuyFeedHandler(t *testing.T) {
	userId := primitive.NewObjectID()
	barnId := primitive.NewObjectID()

	recorder, request := setupRecorderAndRequest(
		httptest.NewRequest("GET", fmt.Sprintf("/barns/%s/buy/feed", barnId), nil),
		userId,
		false,
	)
	mongodb, api := setupMongoMockAndAPI(t)
	request = mux.SetURLVars(request, map[string]string{"barnId": barnId.Hex()})

	api.ServiceCalls.SpendGoldEggs = func(_userId primitive.ObjectID, _amount uint) error {
		return nil
	}

	mongodb.
		EXPECT().
		GetBarn(gomock.Eq(barnId)).
		Return(&barnsvc.Barn{
			Id:            barnId,
			BelongsToUser: userId,
			Feed:          0,
			AutoFeeder:    false,
			DB:            nil,
		}, nil)

	mongodb.
		EXPECT().
		UpdateBarn(gomock.Eq(barnsvc.Barn{
			Id:            barnId,
			BelongsToUser: userId,
			Feed:          10,
			AutoFeeder:    false,
			DB:            api.DB,
		})).
		Return(nil)

	api.buyFeedHandler(recorder, request)
	result := recorder.Result()

	statusCode := result.StatusCode
	if recorder.Result().StatusCode != 200 {
		t.Errorf("Expected StatusCode 200 got %d", statusCode)
	}
}

func TestHandlers_BuyAutoFeederHandler(t *testing.T) {
	userId := primitive.NewObjectID()
	barnId := primitive.NewObjectID()

	recorder, request := setupRecorderAndRequest(
		httptest.NewRequest("GET", fmt.Sprintf("/barns/%s/buy/autoFeeder", barnId), nil),
		userId,
		false,
	)
	mongodb, api := setupMongoMockAndAPI(t)
	request = mux.SetURLVars(request, map[string]string{"barnId": barnId.Hex()})

	api.ServiceCalls.SpendGoldEggs = func(_userId primitive.ObjectID, _amount uint) error {
		return nil
	}

	mongodb.
		EXPECT().
		GetBarn(gomock.Eq(barnId)).
		Return(&barnsvc.Barn{
			Id:            barnId,
			BelongsToUser: userId,
			Feed:          0,
			AutoFeeder:    false,
			DB:            nil,
		}, nil)

	mongodb.
		EXPECT().
		UpdateBarn(gomock.Eq(barnsvc.Barn{
			Id:            barnId,
			BelongsToUser: userId,
			Feed:          0,
			AutoFeeder:    true,
			DB:            nil,
		})).
		Return(nil)

	api.buyAutoFeederHandler(recorder, request)
	result := recorder.Result()

	statusCode := result.StatusCode
	if recorder.Result().StatusCode != 200 {
		t.Errorf("Expected StatusCode 200 got %d", statusCode)
	}
}

func TestHandlers_SpendFeedHandler(t *testing.T) {
	userId := primitive.NewObjectID()
	barnId := primitive.NewObjectID()

	recorder, request := setupRecorderAndRequest(
		httptest.NewRequest(
			"GET",
			"/barns/spendFeed",
			strings.NewReader(
				fmt.Sprintf(`{"barnId": "%s", "amount": %d}`, barnId.Hex(), 1),
			),
		),
		userId,
		true,
	)
	mongodb, api := setupMongoMockAndAPI(t)

	mongodb.
		EXPECT().
		GetBarn(gomock.Eq(barnId)).
		Return(&barnsvc.Barn{
			Id:            barnId,
			BelongsToUser: userId,
			Feed:          1,
			AutoFeeder:    false,
			DB:            nil,
		}, nil)

	mongodb.
		EXPECT().
		UpdateBarn(gomock.Eq(barnsvc.Barn{
			Id:            barnId,
			BelongsToUser: userId,
			Feed:          0,
			AutoFeeder:    false,
			DB:            api.DB,
		})).
		Return(nil)

	api.spendFeedHandler(recorder, request)
	result := recorder.Result()

	statusCode := result.StatusCode
	if recorder.Result().StatusCode != 200 {
		t.Errorf("Expected StatusCode 200 got %d", statusCode)
	}
}

func setupRecorderAndRequest(_request *http.Request, _id primitive.ObjectID, _isService bool) (*httptest.ResponseRecorder, *http.Request) {
	recorder := httptest.NewRecorder()
	token := &util.JwtToken{
		UserId:    _id,
		IsService: _isService,
	}
	ctx := context.WithValue(context.Background(), "jwtToken", token)

	return recorder, _request.WithContext(ctx)
}

func setupMongoMockAndAPI(t *testing.T) (*mocks.MockIMongoDatabase, *Api) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mongodb := mocks.NewMockIMongoDatabase(ctrl)
	api := new(Api)
	api.DB = mongodb

	return mongodb, api
}
