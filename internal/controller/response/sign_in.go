package response

import "time"

type SignIn struct {
	SessionId string    `json:"session_id"`
	ExpiresAt time.Time `json:"expires_at"`
}
