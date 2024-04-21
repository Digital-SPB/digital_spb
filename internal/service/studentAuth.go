package service

import (
	"context"
	"crypto/sha1"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/greenblat17/digital_spb/internal/entity"
	"github.com/greenblat17/digital_spb/internal/repo"
)

const (
	salt = "asdasd"
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

	// return a.repo.CreateAdmin(c, input)
	return s.repo.CreateStudent(ctx, input)

	return 0, nil
}

func (a *StudentAuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
