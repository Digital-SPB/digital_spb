package service

import (
	"context"

	"github.com/greenblat17/digital_spb/internal/entity"
	"github.com/greenblat17/digital_spb/internal/repo"
)

type ExamMarks struct {
	Name string `json:"name"`
	Mark int    `json:"mark"`
}
type StudentAuth interface {
	CreateStudent(ctx context.Context, input entity.Student) (int, error)
	GenerateToken(ctx context.Context, eMail, password string) (string, error)
}

type EducationalDirection interface {
	CreateEducationalDirection(ctx context.Context, education entity.EducatitionalDirection) (int, error)
	CountEducationalDirection(ctx context.Context) (int, error)
}

type Vacancy interface {
	CreateVacancy(ctx context.Context, vacancy entity.Vacancy) (int, error)
	CountVacancy(ctx context.Context) (int, error)
}

type StudyPlan interface {
	GetStudyPlans(ctx context.Context, vacancy entity.Vacancy, examMarks []ExamMarks) ([]entity.EducatitionalDirection, error)
}

type Services struct {
	EducationalDirection EducationalDirection
	Vacancy              Vacancy
	StudyPlan            StudyPlan
	StudentAuth
}

type Applicant interface {
	GetApplicant(ctx context.Context, id int) (entity.Applicant, error)
}

type ServicesDependencies struct {
	Repos *repo.Repositories
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{
		EducationalDirection: NewEducationalDirectionService(deps.Repos.EducationalDirection),
		Vacancy:              NewVacancyService(deps.Repos.Vacancy),
		StudentAuth:          NewStudentAuthService(deps.Repos.StudentAuth),
		StudyPlan:            NewStudyPlanService(deps.Repos.EducationalDirection),
	}
}
