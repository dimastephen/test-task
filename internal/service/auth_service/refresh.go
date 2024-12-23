package auth_service

import (
	"context"
	"errors"
	"os"
	"test-task/internal/api/model"
	auth2 "test-task/internal/repository/model"
	model2 "test-task/internal/service/model"
)

func (a *authService) RefreshJWT(ctx context.Context, tokens model.FormWithTokens) (accessToken string, refreshToken string, err error) {
	ipFromContext := ctx.Value("ip").(string)
	key := os.Getenv("ACCESS_JWT_SIGN")
	if key == "" {
		return "", "", errors.New("failed to get key from .env")
	}

	claims, err := decodeAccessToken(tokens.AccessToken, []byte(key))
	if err != nil {
		return "", "", err
	}

	info, err := a.repo.Get(ctx, auth2.RefreshTokenInfo{RefreshToken: tokens.RefreshToken,
		Ip:   claims.IP,
		Guid: claims.Subject,
	})
	if err != nil {
		return "", "", err
	}

	if ok := ValidateIP(ipFromContext, claims.IP); !ok {
		SendEmailNotification()
	} else if ok = ValidateIP(info.Ip, ipFromContext); !ok {
		SendEmailNotification()
	}

	accessToken, err = getAccessToken(ipFromContext, claims.Subject)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = generateRefreshToken()
	if err != nil {
		return "", "", err
	}

	ids, err := a.repo.Update(ctx, model2.RefreshUpdate{Ip: ipFromContext, NewRefreshToken: refreshToken, Guid: claims.Subject})
	if err != nil {
		return "", "", err
	}
	if ids == 0 {
		return "", "", errors.New("failed to update refresh token")
	}

	return accessToken, refreshToken, nil

}
