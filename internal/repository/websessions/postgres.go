package websessions

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vndg-rdmt/authpod/internal/entity"
)

func NewPostgres(postgres *pgxpool.Pool) Repository {
	return &postgresInstance{
		postgres: postgres,
	}
}

type postgresInstance struct {
	postgres *pgxpool.Pool
}

// Delete implements Repository.
func (p *postgresInstance) Delete(ctx context.Context, sessionId string) error {
	panic("unimplemented")
}

// Get implements Repository.
func (p *postgresInstance) Get(ctx context.Context, sessionId string, result *entity.WebSession) (bool, error) {
	rows := p.postgres.QueryRow(ctx, `
		SELECT
			session_id,
			user_id,
			created_at,
			expires_at
		FROM
			auth.websessions
		WHERE
			session_id = $1
		LIMIT 1;
	`, sessionId)

	if err := rows.Scan(
		&result.SessionId,
		&result.UserId,
		&result.CreatedAt,
		&result.ExpiresAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// Store implements Repository.
func (p *postgresInstance) Create(ctx context.Context, userId int64, expiresAt time.Time) (string, error) {
	var sessId string
	err := p.postgres.QueryRow(ctx, `
		INSERT INTO auth.websessions
		(
			user_id,
			expires_at
		)
		VALUES
		(
			$1,
			$2
		)
		RETURNING session_id;
	`,
		userId,
		expiresAt,
	).Scan(&sessId)

	if err != nil {
		return "", err
	}

	return sessId, nil
}
