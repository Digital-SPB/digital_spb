package pgdb

import (
	"context"

	"github.com/greenblat17/digital_spb/internal/entity"
	"github.com/greenblat17/digital_spb/pkg/postgres"
)

type StudentAuthRepo struct {
	*postgres.Postgres
}

func NewStudentAuthRepo(pg *postgres.Postgres) *StudentAuthRepo {
	return &StudentAuthRepo{pg}
}

// func (r *EducationalDirectionRepo) CreateEducationalDirection(ctx context.Context, education entity.EducatitionalDirection) (int, error) {
// 	var id int
// 	query := "INSERT INTO educational_direction (name, group_name, count_budget, count_contact, price, subject1, subject2, subject3, value1, value2, value3, sum, competive_b, competive_k) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id"

// 	row := r.Pool.QueryRow(ctx, query, education.Name, education.Group, education.CountBudget, education.CountContract,
// 		education.Price, education.Subject1, education.Subject2, education.Subject3, education.Value1, education.Value2, education.Value3, education.Sum, education.CompetiveB, education.CompetiveK)
// 	if err := row.Scan(&id); err != nil {
// 		return 0, err
// 	}

// 	return id, nil
// }

func (r *StudentAuthRepo) CreateStudent(ctx context.Context, input entity.Student) (int, error) {

	return 0, nil
}
