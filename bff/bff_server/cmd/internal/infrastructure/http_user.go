package infrastructure

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type HttpUserGateway struct {
	client *http.Client
	baseUrl string
}

func NewHttpUserGateway(baseUrl string) *HttpUserGateway {
	return &HttpUserGateway{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseUrl: baseUrl,
	}
	
}

func (h *HttpUserGateway) GetUsersByUserID(ctx context.Context, userID string) (*GetUsersByUserIDResponse, error) {
    resp, err := h.client.Get(h.baseUrl + "/users/" + userID)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    var data GetUsersByUserIDResponse
    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        return nil, err
    }
    return &data, nil
}

