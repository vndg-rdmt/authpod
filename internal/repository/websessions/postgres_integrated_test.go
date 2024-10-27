package websessions

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vndg-rdmt/authpod/internal/entity"
)

func TestPost(t *testing.T) {
	postgres, err := pgxpool.New(context.Background(),
		"postgresql://postgres:postgres@127.0.0.1:5432/postgres",
	)
	if err != nil {
		t.Fatalf("cannot connect to postgres: %v", err)
	}

	repo := NewPostgres(postgres)

	sessid, err := repo.Create(context.Background(), 1, time.Now().Add(time.Hour*24))
	if err != nil {
		t.Fatalf("cannot create session: %v", err)
	}
	fmt.Println(sessid)

	var sess entity.WebSession
	ok, err := repo.Get(context.Background(), sessid, &sess)
	if err != nil {
		t.Fatalf("cannot get session: %v", err)
	}

	fmt.Println(ok)
	fmt.Println(sess)
}
