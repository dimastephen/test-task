package service

import (
	"context"
	"test-task/internal/api/model"
)

type AuthService interface {
	GetNewJWT(ctx context.Context) (accessToken string, refreshToken string, err error)
	RefreshJWT(ctx context.Context, tokens model.FormWithTokens) (accesToken string, refreshToken string, err error)
}
