package service

import (
	"context"
	"time"
)

type Token struct {
	Token     string
	ExpiresAt time.Time
}

type Service interface {
	SignIn(ctx context.Context, login, password string) (string, error)
	Ping(ctx context.Context, sessionId string) (int64, error)
	IssueToken(ctx context.Context, result *Token, userId int64) error
	CheckToken(ctx context.Context, token string) (int64, error)
}
