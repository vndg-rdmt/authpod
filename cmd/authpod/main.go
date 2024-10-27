package main

import (
	"context"
	"math/rand"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vndg-rdmt/authpod/internal/controller"
	"github.com/vndg-rdmt/authpod/internal/repository/tokens"
	"github.com/vndg-rdmt/authpod/internal/repository/users"
	"github.com/vndg-rdmt/authpod/internal/repository/websessions"
	"github.com/vndg-rdmt/authpod/internal/service"
	"github.com/vndg-rdmt/authpod/internal/transport"
	"golang.org/x/crypto/bcrypt"
)

const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialBytes = "!@#$%^&*()_+-=[]{}\\|;':\",.<>/?`~"
	numBytes     = "0123456789"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generatePassword(length int, useLetters bool, useSpecial bool, useNum bool) string {
	b := make([]byte, length)
	for i := range b {
		if useLetters {
			b[i] = letterBytes[rand.Intn(len(letterBytes))]
		} else if useSpecial {
			b[i] = specialBytes[rand.Intn(len(specialBytes))]
		} else if useNum {
			b[i] = numBytes[rand.Intn(len(numBytes))]
		}
	}
	return string(b)
}

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func main() {

	postgres, err := pgxpool.New(context.Background(),
		os.Getenv("POSTGRESQL_CONNSTRING"),
	)
	if err != nil {
		panic(err)
	}

	migrate(postgres)

	usersrepo := users.NewPostgres(postgres)
	initroot(usersrepo)

	err = transport.NewHttp(
		controller.NewFiber(
			service.New(
				usersrepo,
				websessions.NewPostgres(postgres),
				tokens.NewPostgres(postgres),
			),
		),
		os.Getenv("LISTEN_ADDR"),
	)
	if err != nil {
		panic(err)
	}
}
