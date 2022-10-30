package auth

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	TokenTTL   = 12 * time.Hour
	salt       = "salt"
	signingKey = "signingKey"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int    `json:"userId"`
	Status string `json:"status"`
}

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func GenerateTokenJWT(userid int, status string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userid,
		status,
	})

	return token.SignedString([]byte(signingKey))
}

type ResultClaims struct {
	UserId int    `json:"userId"`
	Status string `json:"status"`
}

func ParseToken(accessToken string) (*ResultClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*tokenClaims); ok && token.Valid {
		return &ResultClaims{
			UserId: claims.UserId,
			Status: claims.Status,
		}, nil
	}

	return nil, ErrInvalidAccessToken
}

func GenerateTokenRefresh() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}
