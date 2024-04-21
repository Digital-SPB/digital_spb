package pgdb

import (
	"context"

	"github.com/greenblat17/digital_spb/internal/entity"
	"github.com/greenblat17/digital_spb/pkg/postgres"
)

type ApplicantAuthRepo struct {
	*postgres.Postgres
}

func NewApplicantAuthRepo(pg *postgres.Postgres) *ApplicantAuthRepo {
	return &ApplicantAuthRepo{pg}
}

func (r *ApplicantAuthRepo) CreateApplicant(ctx context.Context, input entity.Applicant) (int, error) {
	var id int
	query := "INSERT INTO applicants (name, sure_name, patronymic, email, password_hash, profession) values ($1, $2, $3, $4, $5, $6) RETURNING  id"
	row := r.Pool.QueryRow(ctx, query, input.Name, input.SureName, input.Patronymic, input.EMail, input.Password, input.Profession)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
