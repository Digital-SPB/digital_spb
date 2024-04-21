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
	GetEducationalDirectionForApplicant(ctx context.Context, applicantId int) ([]entity.EducatitionalDirection, error)
}

type Vacancy interface {
	CreateVacancy(ctx context.Context, vacancy entity.Vacancy) (int, error)
	CountVacancy(ctx context.Context) (int, error)
	GetVacancies(ctx context.Context) ([]entity.Vacancy, error)
}

type StudyPlan interface {
	GetStudyPlans(ctx context.Context, id int) ([]entity.EducatitionalDirection, error)
}

type Applicant interface {
	GetApplicant(ctx context.Context, id int) (entity.Applicant, error)
}

type ApplicantAuth interface {
	generatePasswordHash(password string) string
	CreateApplicant(ctx context.Context, input entity.Applicant) (int, error)
}

type Exam interface {
}

type ServicesDependencies struct {
	Repos *repo.Repositories
}

type Services struct {
	EducationalDirection EducationalDirection
	Vacancy              Vacancy
	StudyPlan            StudyPlan
	StudentAuth          StudentAuth
	ApplicantAuth        ApplicantAuth
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{
		EducationalDirection: NewEducationalDirectionService(deps.Repos.EducationalDirection, deps.Repos.Applicant, deps.Repos.Vacancy),
		Vacancy:              NewVacancyService(deps.Repos.Vacancy),
		StudentAuth:          NewStudentAuthService(deps.Repos.StudentAuth),
		StudyPlan:            NewStudyPlanService(deps.Repos.EducationalDirection, deps.Repos.Applicant),
		ApplicantAuth:        NewApplicantAuthService(deps.Repos.ApplicantAuth, deps.Repos.Exam),
	}
}
