package auth

import (
	"github.com/go-chi/chi"
	"test-task/internal/service"
)

type ImplementHandler struct {
	authService service.AuthService
	handler     *chi.Mux
}

func NewImplementHandler(authService service.AuthService) *ImplementHandler {
	mux := chi.NewRouter()
	handler := ImplementHandler{authService: authService, handler: mux}
	mux.Post("/create/{guid}", handler.Create)
	mux.Post("/refresh", handler.Refresh)
	return &handler
}

func (i *ImplementHandler) Handler() *chi.Mux {
	return i.handler
}
