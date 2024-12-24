package auth_service

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

func getAccessToken(ip string, guid string) (accessToken string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, accessClaims{
		IP: ip,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   guid,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})
	signString := os.Getenv("ACCESS_JWT_SIGN")
	signedToken, err := token.SignedString([]byte(signString))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func generateRefreshToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", errors.New("faild to generate refresh token")
	}

	return base64.URLEncoding.EncodeToString(b), nil
}

func SendEmailNotification() {
	log.Println("Отправлен емейл с уведомлением о смене IP")
}

func decodeAccessToken(accessToken string, key []byte) (*accessClaims, error) {
	claims := &accessClaims{}
	parser := jwt.NewParser(jwt.WithoutClaimsValidation())
	token, err := parser.ParseWithClaims(accessToken, claims, func(j *jwt.Token) (interface{}, error) {
		if _, ok := j.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("wrong sign")
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func ValidateIP(ip1 string, ip2 string) bool {
	if ip1 != ip2 {
		return false
	}
	return true
}
