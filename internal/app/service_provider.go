package app

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	auth2 "test-task/internal/api/auth"
	"test-task/internal/config"
	"test-task/internal/config/envs"
	"test-task/internal/repository"
	"test-task/internal/repository/auth"
	"test-task/internal/service"
	"test-task/internal/service/auth_service"
)

type ServiceProvider struct {
	pgConfig   config.PostgresConfig
	httpConfig config.HTTPConfig

	db             *pgxpool.Pool
	authRepo       repository.AuthRepository
	authService    service.AuthService
	implementation *auth2.ImplementHandler
}

func NewServiceProvider() *ServiceProvider {

	return &ServiceProvider{}
}

func (s *ServiceProvider) PGConfig() config.PostgresConfig {
	if s.pgConfig == nil {
		cfg := envs.NewPostgresConfig()
		s.pgConfig = cfg
	}
	return s.pgConfig
}

func (s *ServiceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg := envs.NewHTTPConfig()
		s.httpConfig = cfg
	}
	return s.httpConfig
}

func (s *ServiceProvider) DB(ctx context.Context) *pgxpool.Pool {
	if s.db == nil {
		pool, err := pgxpool.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Connection estabilished")
		s.db = pool
		err = s.db.Ping(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}
	return s.db
}

func (s *ServiceProvider) AuthRepository(ctx context.Context) repository.AuthRepository {
	if s.authRepo == nil {
		s.authRepo = auth.NewAuthRepository(s.DB(ctx))
	}
	return s.authRepo
}

func (s *ServiceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = auth_service.NewAuthService(s.AuthRepository(ctx))
	}
	return s.authService
}

func (s *ServiceProvider) Implementation(ctx context.Context) *auth2.ImplementHandler {
	if s.implementation == nil {
		s.implementation = auth2.NewImplementHandler(s.AuthService(ctx))
	}
	return s.implementation
}
