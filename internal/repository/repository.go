package repository

import (
	"context"
	"test-task/internal/repository/model"
	model2 "test-task/internal/service/model"
)

type AuthRepository interface {
	Create(ctx context.Context, token string) (id int, err error)
	Get(ctx context.Context, info model.RefreshTokenInfo) (*model.RefreshTokenInfo, error)
	Update(ctx context.Context, update model2.RefreshUpdate) (id int, err error)
}
