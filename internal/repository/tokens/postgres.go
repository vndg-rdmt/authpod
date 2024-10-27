package tokens

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
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
func (p *postgresimpl) Create(ctx context.Context, userId int64, expiresAt time.Time) (string, error) {
	var token string

	err := p.postgresql.QueryRow(ctx, `
		INSERT INTO auth.tokens
		(
			user_id,
			expires_at,
			token
		)
		VALUES
		(
			$1,
			$2,
			gen_random_uuid()
		)
		RETURNING token;
	`,
		userId,
		expiresAt,
	).Scan(&token)

	if err != nil {
		return "", err
	}

	return token, nil

}

// GetUserId implements Repository.
func (p *postgresimpl) GetUserId(ctx context.Context, token string) (int64, error) {
	var userId int64
	err := p.postgresql.QueryRow(ctx, `
		SELECT
			user_id
		FROM
			auth.tokens
		WHERE
			token = $1 AND
			expires_at > CURRENT_TIMESTAMP
		LIMIT 1
	`,
		token,
	).Scan(&userId)

	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return 0, ErrNotFound
		}

		return 0, err
	}

	return userId, nil
}
