package chicken_old

//
//func TestGetChickenHandler(t *testing.T) {
//	t.Run("should succeed", func(t *testing.T) {
//		mockController := gomock.NewController(t)
//		defer mockController.Finish()
//
//		chicken := Chicken{
//			ID:            "1",
//			DateOfBirth:   1,
//			EggsLaid:      2,
//			GoldEggsLaid:  3,
//			GoldEggChance: 4,
//			IsAlive:       true,
//		}
//		mockAPI := NewMockAPI(mockController)
//		mockAPI.EXPECT().GetChicken(gomock.Any(), "1").
//			Return(chicken, nil)
//
//		recorder := testHelper(t, mockAPI, http.MethodGet, "/chickens/1", nil)
//		assert.Equal(t, http.StatusOK, recorder.Code)
//
//		var resultChicken Chicken
//		if err := json.Unmarshal(recorder.Body.Bytes(), &resultChicken); err != nil {
//			t.Fatal(err)
//		}
//		assert.Equal(t, chicken, resultChicken)
//	})
//
//	t.Run("should 500 on engine error", func(t *testing.T) {
//		mockController := gomock.NewController(t)
//		defer mockController.Finish()
//
//		mockAPI := NewMockAPI(mockController)
//		mockAPI.EXPECT().GetChicken(gomock.Any(), "1").
//			Return(Chicken{}, errors.New("fake error"))
//
//		recorder := testHelper(t, mockAPI, http.MethodGet, "/chickens/1", nil)
//
//		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
//	})
//
//}
//
//func TestNewChickenHandler(t *testing.T) {
//	t.Run("should succeed", func(t *testing.T) {
//		mockController := gomock.NewController(t)
//		defer mockController.Finish()
//
//		mockAPI := NewMockAPI(mockController)
//		mockAPI.EXPECT().NewChicken(gomock.Any(), "1").Return("1", nil)
//
//		recorder := testHelper(t, mockAPI, http.MethodPost, "/chickens/new",
//			NewChickenRequest{FarmID: "1"},
//		)
//		assert.Equal(t, http.StatusOK, recorder.Code)
//
//		var result NewChickenResult
//		if err := json.Unmarshal(recorder.Body.Bytes(), &result); err != nil {
//			t.Fatal(err)
//		}
//		assert.Equal(t, NewChickenResult{ChickenID: "1"}, result)
//	})
//
//	t.Run("should 400 on missing body", func(t *testing.T) {
//		recorder := testHelper(t, nil, http.MethodPost, "/chickens/new", "")
//		assert.Equal(t, http.StatusBadRequest, recorder.Code)
//	})
//
//	t.Run("should 500 on engine error", func(t *testing.T) {
//		mockController := gomock.NewController(t)
//		defer mockController.Finish()
//
//		mockAPI := NewMockAPI(mockController)
//		mockAPI.EXPECT().NewChicken(gomock.Any(), "1").
//			Return("", errors.New("fake error"))
//
//		recorder := testHelper(t, mockAPI, http.MethodPost, "/chickens/new",
//			NewChickenRequest{FarmID: "1"},
//		)
//		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
//	})
//
//}
//
//func testHelper(
//	t *testing.T, engine API, method, url string, body any,
//) *httptest.ResponseRecorder {
//	var server = RESTServer{
//		Root:   mux.NewRouter(),
//		Engine: engine,
//	}
//	server.SetupHandlers()
//
//	var reader io.Reader
//	if body != nil {
//		body, err := json.Marshal(NewChickenRequest{FarmID: "1"})
//		if err != nil {
//			t.Fatal(err)
//		}
//
//		reader = bytes.NewReader(body)
//	}
//
//	req, err := http.NewRequest(method, url, reader)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	recorder := httptest.NewRecorder()
//	server.Root.ServeHTTP(recorder, req)
//
//	return recorder
//}
