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
}

type StudentAuth interface {
	CreateStudent(ctx context.Context, input entity.Student) (int, error)
}

type Repositories struct {
	EducationalDirection
	StudentAuth
}

func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		EducationalDirection: pgdb.NewEducationalDirectionRepo(pg),
		StudentAuth:          pgdb.NewStudentAuthRepo(pg),
	}
}
