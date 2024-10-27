package users

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/vndg-rdmt/authpod/internal/entity"
)

type PostgresConn interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

func NewPostgres(conn PostgresConn) Repository {
	return &postgresimpl{
		postgresql: conn,
	}
}

type postgresimpl struct {
	postgresql PostgresConn
}

// Create implements Repository.
func (p *postgresimpl) Create(
	ctx context.Context,
	login string,
	passwordHash string,
) (int64, error) {
	var userId int64
	err := p.postgresql.QueryRow(ctx, `
		INSERT INTO auth.users
		(
			login,
			password_hash
		)
		VALUES
		(
			$1,
			$2
		)
		RETURNING id;
	`,
		login,
		passwordHash,
	).Scan(&userId)

	if err != nil {
		return 0, err
	}

	return userId, nil

}

// GetByCredentials implements Repository.
func (p *postgresimpl) GetByLogin(
	ctx context.Context,
	result *entity.User,
	login string,
) error {
	err := p.postgresql.QueryRow(ctx, `
		SELECT
			id,
			created_at,
			login,
			password_hash
		FROM
			auth.users
		WHERE
			login = $1
		LIMIT 1
	`,
		login,
	).Scan(
		&result.Id,
		&result.CreatedAt,
		&result.Login,
		&result.PassworHash,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrNotFound
		}
		return err
	}

	return nil
}
