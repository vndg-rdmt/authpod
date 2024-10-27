package users

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vndg-rdmt/authpod/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func TestPostgres(t *testing.T) {
	postgres, err := pgxpool.New(context.Background(),
		"postgresql://postgres:postgres@127.0.0.1:5432/postgres",
	)
	if err != nil {
		t.Fatalf("cannot connect to postgres: %v", err)
	}

	repo := NewPostgres(postgres)

	userid, err := repo.Create(context.Background(), "123452222", hashPassword("12345"))
	if err != nil {
		t.Fatalf("cannot create user: %v", err)
	}
	fmt.Println(userid)

	var user entity.User
	if err := repo.GetByLogin(context.Background(), &user, "123452222"); err != nil {
		t.Fatalf("cannot get user: %v", err)
	}
	fmt.Println(user)
}
