package infrastructure

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type HttpOrderGateway struct {
	client *http.Client
	baseUrl string
}

func NewHttpOrderGateway(baseUrl string) *HttpOrderGateway {
	return &HttpOrderGateway{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseUrl: baseUrl,
	}
	
}

func (h *HttpOrderGateway) GetOrdersByUserID(ctx context.Context, userID string) (*GetOrdersByUserIDResponse, error) {
	resp, err := h.client.Get(h.baseUrl + "/orders/" + userID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data GetOrdersByUserIDResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	
	return &data, nil
}


