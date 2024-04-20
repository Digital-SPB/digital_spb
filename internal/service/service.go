package service

import "github.com/greenblat17/digital_spb/internal/repo"

type Services struct {
}

type ServicesDependencies struct {
	Repos *repo.Repositories
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{}
}
