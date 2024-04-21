package pgdb

import (
	"context"

	"github.com/greenblat17/digital_spb/internal/entity"
	"github.com/greenblat17/digital_spb/pkg/postgres"
)

type ApplicantRepo struct {
	*postgres.Postgres
}

func NewApplicantRepo(pg *postgres.Postgres) *ApplicantRepo {
	return &ApplicantRepo{pg}
}

func (p *ApplicantRepo) GetApplicant(ctx context.Context, id int) (entity.Applicant, error) {
	query := "SELECT * FROM applicants WHERE id = $1"
	row := p.Pool.QueryRow(ctx, query, id)
	var applicant entity.Applicant
	if err := row.Scan(&applicant); err != nil {
		return entity.Applicant{}, err
	}

	return applicant, nil
}
