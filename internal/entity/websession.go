package entity

import "time"

type WebSession struct {
	Id          string    `json:"id"`
	UserId      string    `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	ExpiresAt   time.Time `json:"expires_at"`
	Fingerprint string    `json:"fingerprint"`
}
