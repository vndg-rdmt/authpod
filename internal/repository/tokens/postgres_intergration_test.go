package tokens

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestPostgreSQL(t *testing.T) {
	postgres, err := pgxpool.New(context.Background(),
		"postgresql://postgres:postgres@127.0.0.1:5432/postgres",
	)
	if err != nil {
		t.Fatalf("cannot connect to postgres: %v", err)
	}

	repo := NewPostgres(postgres)

	token, err := repo.Create(context.Background(), 1, time.Now().Add(time.Hour))
	if err != nil {
		t.Fatalf("cannot create new token: %v", err)
	}

	fmt.Println(token)

	userId, err := repo.GetUserId(context.Background(), token)
	if err != nil {
		t.Fatalf("cannot create new token: %v", err)
	}

	t.Log(userId)
}
