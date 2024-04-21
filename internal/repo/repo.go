package repo

import (
	"context"

	"github.com/greenblat17/digital_spb/internal/entity"
	"github.com/greenblat17/digital_spb/internal/repo/pgdb"
	"github.com/greenblat17/digital_spb/pkg/postgres"
)

type EducationalDirection interface {
	CreateEducationalDirection(ctx context.Context, education entity.EducatitionalDirection) (int, error)
	CountEducationalDirection(ctx context.Context) (int, error)
	GetEducationalDirections(ctx context.Context) ([]entity.EducatitionalDirection, error)
}

type StudentAuth interface {
	CreateStudent(ctx context.Context, input entity.Student) (int, error)
	GetStudent(ctx context.Context, eMail, paswwrord string) (entity.Student, error)
}

type Vacancy interface {
	CreateVacancy(ctx context.Context, vacancy entity.Vacancy) (int, error)
	CountVacancy(ctx context.Context) (int, error)
	GetVacanciesByEducationId(ctx context.Context, educationId int) ([]entity.Vacancy, error)
}

type Applicant interface {
	GetApplicant(ctx context.Context, id int) (entity.Applicant, error)
	GetExam(ctx context.Context, applicantId int) ([]entity.Exam, error)
}

type Repositories struct {
	EducationalDirection
	StudentAuth
	Vacancy
	Applicant
}

func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		EducationalDirection: pgdb.NewEducationalDirectionRepo(pg),
		StudentAuth:          pgdb.NewStudentAuthRepo(pg),
		Vacancy:              pgdb.NewVacancyRepo(pg),
		Applicant:            pgdb.NewApplicantRepo(pg),
	}
}
