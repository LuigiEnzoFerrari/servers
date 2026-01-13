package http_client

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"github.com/LuigiEnzoFerrari/servers/bff/dashboard_server/cmd/internal/domain"
)

type HttpOrderGateway struct {
	client  *http.Client
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

func (h *HttpOrderGateway) GetOrdersByUserID(ctx context.Context, userID string) ([]domain.ExternalOrder, error) {
	resp, err := h.client.Get(h.baseUrl + "/" + userID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data GetOrdersByUserIDResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return MapGetOrdersByUserIDResponseToExternalOrder(data), nil
}
