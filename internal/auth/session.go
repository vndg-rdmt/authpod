package auth

import (
	"context"

	"github.com/vndg-rdmt/authpod/internal/entity"
	"github.com/vndg-rdmt/authpod/internal/repository/websessions"
)

const (
	MethodWebSession = "Web-Session"
)

func NewSessionsMethod(repo websessions.Repository) AuthenticationMethod {
	return &sessionsMethod{
		websessions: repo,
	}
}

type sessionsMethod struct {
	websessions websessions.Repository
}

// Name implements AuthenticationMethod.
func (s *sessionsMethod) Name() string {
	return MethodWebSession
}

// Authenticate implements AuthenticationMethod.
func (s *sessionsMethod) Authenticate(ctx context.Context, sess *entity.User, secret string) (bool, error) {
	var session entity.WebSession

	ok, err := s.websessions.Get(ctx, secret, &session)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}

	sess.Id = session.UserId
	return false, nil
}
