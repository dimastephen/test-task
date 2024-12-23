package auth

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"log"
	"net"
	"net/http"
	"test-task/internal/api/model"
)

func (i *ImplementHandler) Create(w http.ResponseWriter, r *http.Request) {
	guid := chi.URLParam(r, "guid")
	userIp, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
	}
	ctx := context.Background()
	ctx = addRequestToContext(ctx, guid, userIp)
	accessToken, refreshToken, err := i.authService.GetNewJWT(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
	}
	jsonResponse := model.FormWithTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expires:      720,
		Type:         "Bearer",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(jsonResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func addRequestToContext(ctx context.Context, guid string, ip string) context.Context {
	ctx = context.WithValue(ctx, "guid", guid)
	ctx = context.WithValue(ctx, "ip", ip)
	return ctx
}
