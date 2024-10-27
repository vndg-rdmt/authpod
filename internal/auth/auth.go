package auth

import (
	"context"

	"github.com/vndg-rdmt/authpod/internal/entity"
)

func NewAuthentication(methods ...AuthenticationMethod) Authentication {

	m := &authentication{
		methods: make(map[string]AuthenticationMethod),
	}

	for i := 0; i < len(methods); i++ {
		m.methods[methods[i].Name()] = methods[i]
	}

	return m
}

type authentication struct {
	methods map[string]AuthenticationMethod
}

// Authenticate implements Authentication.
func (a *authentication) Authenticate(ctx context.Context, sess *entity.User, methodName string, secret string) (bool, error) {
	method, ok := a.methods[methodName]
	if !ok {
		return false, nil
	}
	return method.Authenticate(ctx, sess, secret)
}
