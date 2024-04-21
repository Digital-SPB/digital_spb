package service

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/greenblat17/digital_spb/internal/entity"
	"github.com/greenblat17/digital_spb/internal/repo"
)

const (
	salt = "asdasd"
	TokenTTL = 12 * time.Hour
	signingKey = "123iuasd2rke;aldj"
)

type StudentAuthService struct {
	repo repo.StudentAuth
}

type tokenClaims struct {
	jwt.RegisteredClaims
	Id int `json:"id"`
}

func NewStudentAuthService(repo repo.StudentAuth) *StudentAuthService {
	return &StudentAuthService{repo: repo}
}

func (s *StudentAuthService) CreateStudent(ctx context.Context, input entity.Student) (int, error) {
	input.Password = s.generatePasswordHash(input.Password)

	return s.repo.CreateStudent(ctx, input)
}

func (a *StudentAuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (a *StudentAuthService) GenerateToken(ctx context.Context, eMail, password string) (string, error) {
	admin, err := a.repo.GetStudent(ctx, eMail, a.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Id: admin.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (a *StudentAuthService) Parsetoken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, nil
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.Id, nil
}
