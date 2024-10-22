package websessions

import (
	"context"

	"github.com/vndg-rdmt/authpod/internal/entity"
)

type Repository interface {
	Get(ctx context.Context, sessionId string, result *entity.WebSession) (bool, error)
	Store(ctx context.Context, session *entity.WebSession) error
	Delete(ctx context.Context, sessionId string) error
}
