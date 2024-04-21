package service

import (
	"context"

	"github.com/greenblat17/digital_spb/internal/entity"
	"github.com/greenblat17/digital_spb/internal/repo"
)

type ApplicantService struct {
	ApplicantRepo repo.Applicant
}

func (s *ApplicantService) GetApplicant(ctx context.Context, id int) (entity.Applicant, error) {
	return s.ApplicantRepo.GetApplicant(ctx, id)
}
