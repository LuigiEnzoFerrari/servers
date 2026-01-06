package infrastructure

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/LuigiEnzoFerrari/servers/bff/bff_server/cmd/internal/domain"
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

func (h *HttpOrderGateway) GetOrdersByUserID(ctx context.Context, userID string) (*domain.GetOrdersByUserIDResponse, error) {
	resp, err := http.Get(h.baseUrl + "/orders/" + userID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data domain.GetOrdersByUserIDResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	
	return &data, nil
}


