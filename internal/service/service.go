package service

import (
	"context"

	"github.com/greenblat17/digital_spb/internal/entity"
	"github.com/greenblat17/digital_spb/internal/repo"
)

type StudentAuth interface {
	CreateStudent(ctx context.Context, input entity.Student) (int, error)
}

type Services struct {
	StudentAuth
}

type ServicesDependencies struct {
	Repos *repo.Repositories
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{}
}
