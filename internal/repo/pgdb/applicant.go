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

func (p *ApplicantRepo) GetExam(ctx context.Context, applicantId int) ([]entity.Exam, error) {
	query := "SELECT * FROM applicants WHERE applicant_id = $1"
	rows, err := p.Pool.Query(ctx, query, applicantId)
	if err != nil {
		return nil, err
	}

	var exams []entity.Exam
	for rows.Next() {
		var exam entity.Exam
		if err := rows.Scan(&exam.Id, &exam.ExamName, &exam.ExamMark, &exam.ApplicantId); err != nil {
			return nil, err
		}
		exams = append(exams, exam)
	}

	return exams, nil
}
