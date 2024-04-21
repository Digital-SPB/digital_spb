package service

import (
	"context"

	"github.com/greenblat17/digital_spb/internal/entity"
	"github.com/greenblat17/digital_spb/internal/repo"
)

type StudyPlanService struct {
	educationalDirectionRepo repo.EducationalDirection
	applicantRepo            repo.Applicant
}

func NewStudyPlanService(educationalDirectionRepo repo.EducationalDirection, applicantRepo repo.Applicant) *StudyPlanService {
	return &StudyPlanService{
		educationalDirectionRepo: educationalDirectionRepo,
		applicantRepo:            applicantRepo,
	}
}

func (s *StudyPlanService) GetStudyPlans(ctx context.Context, id int) ([]entity.EducatitionalDirection, error) {
	return []entity.EducatitionalDirection{}, nil
}
