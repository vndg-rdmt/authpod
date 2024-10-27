package tokens

import (
	"context"
	"errors"
	"time"
)

var (
	ErrNotFound = errors.New("not found")
)

type Repository interface {
	Create(ctx context.Context, userId int64, expiresAt time.Time) (string, error)
	GetUserId(ctx context.Context, token string) (int64, error)
}
