package service

import (
	"context"

	"github.com/greenblat17/digital_spb/internal/entity"
	"github.com/greenblat17/digital_spb/internal/repo"
)

type VacancyService struct {
	vacancyRepo repo.Vacancy
}

func NewVacancyService(vacancyRepo repo.Vacancy) *VacancyService {
	return &VacancyService{vacancyRepo: vacancyRepo}
}

func (s *VacancyService) CreateVacancy(ctx context.Context, vacancy entity.Vacancy) (int, error) {
	return s.vacancyRepo.CreateVacancy(ctx, vacancy)
}

func (s *VacancyService) CountVacancy(ctx context.Context) (int, error) {
	return s.vacancyRepo.CountVacancy(ctx)
}
