package chicken_old

import (
	"bytes"
	"chicken-farmer/backend/internal"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RESTChickenClient struct {
	httpClient *http.Client
	baseURL    string
}

func ProvideRESTChickenClient(baseURL string) API {
	return &RESTChickenClient{
		httpClient: internal.NewHTTPClient(),
		baseURL:    baseURL,
	}
}

func (client RESTChickenClient) GetChicken(
	ctx context.Context, chickenID string,
) (Chicken, error) {
	var result Chicken
	if err := client.doRequest(
		ctx, http.MethodGet, fmt.Sprintf("/chickens/%s", chickenID), nil, &result,
	); err != nil {
		return result, err
	}

	return result, nil
}

func (client RESTChickenClient) NewChicken(
	ctx context.Context, farmID string,
) (string, error) {
	chickenBytes, err := json.Marshal(NewChickenRequest{
		FarmID: farmID,
	})
	if err != nil {
		return "", err
	}

	var result NewChickenResult
	if err := client.doRequest(
		ctx, http.MethodPost, "/chickens/new", chickenBytes, &result,
	); err != nil {
		return "", err
	}

	return result.ChickenID, nil
}

func (client RESTChickenClient) FeedChicken(
	ctx context.Context, chickenID string,
) error {
	return client.doRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("/chickens/%s/feed", chickenID),
		nil,
		nil,
	)
}

func (client RESTChickenClient) doRequest(
	ctx context.Context, method, urlPath string, body []byte, result any,
) error {
	request, err := http.NewRequestWithContext(
		ctx,
		method,
		fmt.Sprintf("%s%s", client.baseURL, urlPath), bytes.NewReader(body),
	)
	if err != nil {
		return err
	}

	response, err := client.httpClient.Do(request)
	if err != nil {
		return err
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("[%d]: %s", response.StatusCode, responseBody)
	}

	if result != nil {
		if err := json.Unmarshal(responseBody, result); err != nil {
			return err
		}
	}

	return nil
}
