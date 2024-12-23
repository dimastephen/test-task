package app

import (
	"context"
	"log"
	"net/http"
	"test-task/internal/api/auth"
	"test-task/internal/config"
	"time"
)

type app struct {
	provider *ServiceProvider
	server   *http.Server
	authImpl auth.ImplementHandler
}

func NewApp(ctx context.Context, configPath string) (*app, error) {
	a := &app{}
	err := a.InitVars(ctx, configPath)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *app) InitVars(ctx context.Context, configPath string) error {
	err := a.initConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	inits := []func(ctx context.Context) error{
		a.initServiceProvider,
		a.initServer,
	}
	for _, fun := range inits {
		err := fun(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *app) initConfig(configPath string) error {
	err := config.Load(configPath)
	if err != nil {
		return err
	}
	return nil
}

func (a *app) initServiceProvider(_ context.Context) error {
	a.provider = NewServiceProvider()
	return nil
}

func (a *app) initServer(ctx context.Context) error {
	a.server = &http.Server{
		Addr:         a.provider.HTTPConfig().Address(),
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		Handler:      a.provider.Implementation(ctx).Handler(),
	}
	return nil
}

func (a *app) RunHttp() error {
	err := a.server.ListenAndServe()
	return err
}
