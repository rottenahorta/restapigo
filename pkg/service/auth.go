package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	todo "github.com/rottenahorta/restapi101"
	"github.com/rottenahorta/restapi101/pkg/repo"
)

const (
	salt     = "jhiufhuihfleu8e9fhe"
	jwtKey   = "hiuhfieulhuinehfiu&^&*Tgk^&"
	tokenTTL = 12 * time.Hour
)

type AuthService struct {
	repo repo.Auth
}

type tokenClaims struct {
	jwt.StandardClaims
	UID int `json:"user_id"`
}

func NewAuthService(r repo.Auth) *AuthService {
	return &AuthService{repo: r}
}

func (s *AuthService) CreateUser(u todo.User) (int, error) {
	u.Password = generatePasswordHash(u.Password)
	return s.repo.CreateUser(u)
}

func (s *AuthService) GenerateToken(un, p string) (string, error) { // public go after public method, then private - clean code
	u, err := s.repo.GetUserCred(un, generatePasswordHash(p))
	if err != nil {
		return "", err
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt: time.Now().Unix()},
		u.Id,
	})
	return t.SignedString([]byte(jwtKey))
}

func (s *AuthService) ParseJWT(token string) (int, error) {
	t, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok { // checkin if signin is HMAC
			return nil, errors.New("invalid signin method")
		}
		return []byte(jwtKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := t.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("claims rnt of local type  *tokenClaims")
	}
	return claims.UID, nil
}

func generatePasswordHash(p string) string {
	h := sha1.New()
	h.Write(([]byte(p)))
	return fmt.Sprintf("%x", h.Sum([]byte(salt)))
}
