package auth

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"test-task/internal/api/model"
)

func (i *ImplementHandler) Refresh(w http.ResponseWriter, r *http.Request) {

	req := model.FormWithTokens{}

	userIp, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		http.Error(w, "failed to get ip", http.StatusInternalServerError)
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, "ip", userIp)

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		http.Error(w, "wrong request body", http.StatusInternalServerError)
		log.Fatal(err)
	}
	accessToken, refreshToken, err := i.authService.RefreshJWT(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	response := model.FormWithTokens{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
		Expires:      720,
		Type:         "Bearer",
	}
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
