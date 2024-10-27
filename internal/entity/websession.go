package entity

import "time"

type WebSession struct {
	SessionId string    `json:"session_id"`
	UserId    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}
