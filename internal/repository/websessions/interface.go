package websessions

import (
	"context"
	"time"

	"github.com/vndg-rdmt/authpod/internal/entity"
)

type Repository interface {
	Get(ctx context.Context, sessionId string, result *entity.WebSession) (bool, error)
	Create(ctx context.Context, userId int64, expiresAt time.Time) (string, error)
	Delete(ctx context.Context, sessionId string) error
}
