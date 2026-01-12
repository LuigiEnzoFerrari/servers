package http_client

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/LuigiEnzoFerrari/servers/bff/dashboard_server/cmd/internal/domain"
)

type HttpUserGateway struct {
	client  *http.Client
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

func (h *HttpUserGateway) GetUserByUserID(ctx context.Context, userID string) (*domain.ExternalUser, error) {
	resp, err := h.client.Get(h.baseUrl + "/users/" + userID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data GetUserByUserIDResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return MapGetUserByUserIDResponseToExternalUser(data), nil
}
