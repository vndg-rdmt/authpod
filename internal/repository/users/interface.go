package users

import (
	"context"
	"errors"

	"github.com/vndg-rdmt/authpod/internal/entity"
)

var (
	ErrNotFound = errors.New("not found")
)

type Repository interface {
	GetByLogin(
		ctx context.Context,
		result *entity.User,
		login string,
	) error

	Create(
		ctx context.Context,
		login string,
		passwordHash string,
	) (int64, error)
}
