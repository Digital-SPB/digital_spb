package pgdb

import (
	"context"
	"fmt"

	"github.com/greenblat17/digital_spb/internal/entity"
	"github.com/greenblat17/digital_spb/pkg/postgres"
)

type ExamRepo struct {
	*postgres.Postgres
}

func NewExamRepo(pg *postgres.Postgres) *ExamRepo {
	return &ExamRepo{pg}
}

func (r *ExamRepo) AddExam(ctx context.Context, input entity.Exam) (int, error){
	var id int
	fmt.Println("11")
	query := "INSERT INTO exams (applicant_id, exam_name, exam_mark) values ($1, $2, $3) RETURNING  id"
	fmt.Println("22")
	row := r.Pool.QueryRow(ctx, query, input.ApplicantId, input.ExamName, input.ExamMark)
	fmt.Println("33")
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	fmt.Println("44")
	return id, nil
}
