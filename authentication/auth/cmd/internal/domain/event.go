package domain

import "time"

type Event struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Payload   any       `json:"payload"`
	OccurredAt time.Time `json:"occurred_at"`
}