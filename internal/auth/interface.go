package auth

import (
	"context"

	"github.com/vndg-rdmt/authpod/internal/entity"
)

type Authentication interface {
	Authenticate(ctx context.Context, sess *entity.User, method, secret string) (bool, error)
}

type AuthenticationMethod interface {
	Name() string
	Authenticate(ctx context.Context, sess *entity.User, secret string) (bool, error)
}
