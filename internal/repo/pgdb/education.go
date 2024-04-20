package pgdb

import (
	"context"

	"github.com/greenblat17/digital_spb/internal/entity"
	"github.com/greenblat17/digital_spb/pkg/postgres"
)

type EducationalDirectionRepo struct {
	*postgres.Postgres
}

func NewEducationalDirectionRepo(pg *postgres.Postgres) *EducationalDirectionRepo {
	return &EducationalDirectionRepo{pg}
}

func (r *EducationalDirectionRepo) CreateEducationalDirection(ctx context.Context, education entity.EducatitionalDirection) (int, error) {
	var id int
	query := "INSERT INTO educational_direction (name, group_name, count_budget, count_contact, price, subject1, subject2, subject3, value1, value2, value3, sum, competive_b, competive_k) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id"

	row := r.Pool.QueryRow(ctx, query, education.Name, education.Group, education.CountBudget, education.CountContract,
		education.Price, education.Subject1, education.Subject2, education.Subject3, education.Value1, education.Value2, education.Value3, education.Sum, education.CompetiveB, education.CompetiveK)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *EducationalDirectionRepo) CountEducationalDirection(ctx context.Context) (int, error) {
	query := "SELECT count(*) FROM educational_direction"
	var count int
	row := r.Pool.QueryRow(ctx, query)
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}
