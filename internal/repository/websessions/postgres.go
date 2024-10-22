package websessions

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vndg-rdmt/authpod/internal/entity"
)

func NewPostgreSQL(postgres *pgxpool.Pool) Repository {
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
			id,
			user_id,
			created_at,
			expires_at,
			fingerprint
		FROM
			auth.web_sessions
		WHERE
			id = $1
		LIMIT 1;
	`, sessionId)

	if err := rows.Scan(
		&result.Id,
		&result.UserId,
		&result.CreatedAt,
		&result.ExpiresAt,
		&result.Fingerprint,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// Store implements Repository.
func (p *postgresInstance) Store(ctx context.Context, session *entity.WebSession) error {
	_, err := p.postgres.Exec(ctx, `
		INSERT INTO auth.web_sessions
		(
			id,
			user_id,
			created_at,
			expires_at,
			fingerprint
		)
		VALUES
		(
			$1,
			$2,
			$3,
			$4
		);
	`,
		session.Id,
		session.UserId,
		session.CreatedAt,
		session.ExpiresAt,
		session.Fingerprint,
	)
	return err
}
