package service

import (
	"context"

	"github.com/greenblat17/digital_spb/internal/entity"
	"github.com/greenblat17/digital_spb/internal/repo"
)

type EducationalDirectionService struct {
	educationalDirectionRepo repo.EducationalDirection
}

func NewEducationalDirectionService(educationalDirectionRepo repo.EducationalDirection) *EducationalDirectionService {
	return &EducationalDirectionService{educationalDirectionRepo: educationalDirectionRepo}
}

func (s *EducationalDirectionService) CreateEducationalDirection(ctx context.Context, education entity.EducatitionalDirection) (int, error) {
	return s.educationalDirectionRepo.CreateEducationalDirection(ctx, education)
}

func (s *EducationalDirectionService) CountEducationalDirection(ctx context.Context) (int, error) {
	return s.educationalDirectionRepo.CountEducationalDirection(ctx)
}
