package auth_service

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

type accessClaims struct {
	jwt.RegisteredClaims
	IP string
}

func (a *authService) GetNewJWT(ctx context.Context) (accessToken string, refreshToken string, err error) {
	ip := ctx.Value("ip").(string)
	guid := ctx.Value("guid").(string)

	if ip == "" || guid == "" {
		return "", "", errors.New("failed to pars ip or guid")
	}
	accessToken, err = getAccessToken(ip, guid)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = generateRefreshToken()
	if err != nil {
		return "", "", err
	}
	id, err := a.repo.Create(ctx, refreshToken)
	if err != nil || id == 0 {
		return "", "", err
	}

	return accessToken, refreshToken, nil

}
