package service

import (
	"context"

	"github.com/greenblat17/digital_spb/internal/entity"
	"github.com/greenblat17/digital_spb/internal/repo"
)

type StudyPlanService struct {
	educationalDirectionRepo repo.EducationalDirection
}

func NewStudyPlanService(educationalDirectionRepo repo.EducationalDirection) *StudyPlanService {
	return &StudyPlanService{educationalDirectionRepo: educationalDirectionRepo}
}

func (s *StudyPlanService) GetStudyPlans(ctx context.Context, vacancy entity.Vacancy, examMarks []ExamMarks) ([]entity.EducatitionalDirection, error) {
	return []entity.EducatitionalDirection{}, nil
}
