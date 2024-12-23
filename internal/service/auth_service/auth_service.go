package auth_service

import (
	"test-task/internal/repository"
	"test-task/internal/service"
)

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) service.AuthService {
	return &authService{
		repo: repo,
	}
}
