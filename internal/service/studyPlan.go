package service

import (
	"context"

	"github.com/greenblat17/digital_spb/internal/entity"
	"github.com/greenblat17/digital_spb/internal/repo"
)

type StudyPlanService struct {
	studyPlan repo.StudyPlan
}

func NewStudyPlanService(studyPlan repo.StudyPlan) *StudyPlanService {
	return &StudyPlanService{studyPlan: studyPlan}
}

func (s *StudyPlanService) GetStudyPlans(ctx context.Context, vacancy entity.Vacancy) (int, error) {
	return 0, nil
}
