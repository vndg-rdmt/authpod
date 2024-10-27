package service

import (
	"context"
	"fmt"
	"time"

	"github.com/vndg-rdmt/authpod/internal/entity"
	"github.com/vndg-rdmt/authpod/internal/repository/tokens"
	"github.com/vndg-rdmt/authpod/internal/repository/users"
	"github.com/vndg-rdmt/authpod/internal/repository/websessions"
	"golang.org/x/crypto/bcrypt"
)

func New(
	users users.Repository,
	websession websessions.Repository,
	tokens tokens.Repository,
) Service {
	return &serviceimpl{
		users:      users,
		websession: websession,
		tokens:     tokens,
	}
}

type serviceimpl struct {
	users      users.Repository
	websession websessions.Repository
	tokens     tokens.Repository
}

// SignIn implements Service.
func (s *serviceimpl) SignIn(ctx context.Context, login string, password string) (string, error) {

	var user entity.User
	var sessionId string
	var err error

	if err = s.users.GetByLogin(context.Background(), &user, login); err != nil {
		if err == users.ErrNotFound {
			return "", ErrNotFound
		}
		return "", fmt.Errorf("cannot get user by login: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PassworHash), []byte(password)); err != nil {
		return "", ErrNotFound
	}

	if sessionId, err = s.websession.Create(ctx, user.Id, time.Now().Add(time.Hour*24)); err != nil {
		return "", fmt.Errorf("cannot create new session: %w", err)
	}

	return sessionId, nil
}

// SignIn implements Service.
func (s *serviceimpl) Ping(ctx context.Context, sessionId string) (int64, error) {
	var sess entity.WebSession

	ok, err := s.websession.Get(ctx, sessionId, &sess)
	if err != nil {
		return 0, err
	}

	if ok {
		return sess.UserId, nil
	}

	return 0, ErrNotFound
}

func (s *serviceimpl) IssueToken(ctx context.Context, result *Token, userId int64) error {

	expiresAt := time.Now().Add(time.Hour * 24)

	sessId, err := s.tokens.Create(ctx, userId, expiresAt)
	if err != nil {
		return err
	}

	result.ExpiresAt = expiresAt
	result.Token = sessId

	return nil
}

func (s *serviceimpl) CheckToken(ctx context.Context, token string) (int64, error) {
	userId, err := s.tokens.GetUserId(ctx, token)
	if err != nil {
		if err == tokens.ErrNotFound {
			return 0, ErrNotFound
		}
		return 0, err
	}

	return userId, nil
}
