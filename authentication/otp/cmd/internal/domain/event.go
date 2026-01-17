package domain

import (
	"context"
	"time"
	"encoding/json"
)


type PasswordForgotEvent struct {
	Username string `json:"username"`
}


type Event struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	TraceID   string    `json:"trace_id"`
	Payload   json.RawMessage       `json:"payload"`
	OccurredAt time.Time `json:"occurred_at"`
}

type EventHandler func(ctx context.Context, event Event) error