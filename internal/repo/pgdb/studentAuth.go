package pgdb

import (
	"context"
	"errors"

	"github.com/greenblat17/digital_spb/internal/entity"
	"github.com/greenblat17/digital_spb/pkg/postgres"
	"github.com/jackc/pgx/v5"
)

type StudentAuthRepo struct {
	*postgres.Postgres
}

func NewStudentAuthRepo(pg *postgres.Postgres) *StudentAuthRepo {
	return &StudentAuthRepo{pg}
}

func (r *StudentAuthRepo) CreateStudent(ctx context.Context, input entity.Student) (int, error) {
	var id int
	query := "INSERT INTO students (name, sure_name, patronymic, email, password_hash, university, direction, group_number) values ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING  id"
	row := r.Pool.QueryRow(ctx, query, input.Name, input.SureName, input.Patronymic, input.EMail, input.Password, input.University, input.Direction, input.Group)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *StudentAuthRepo) GetStudent(ctx context.Context, eMail, paswwrord string) (entity.Student, error) {
	var student entity.Student
	query := "SELECT * FROM students WHERE email = $1 AND password_hash = $2"
	row := r.Pool.QueryRow(ctx, query, eMail, paswwrord)
	if err := row.Scan(&student.Id, &student.Name, &student.SureName, &student.Patronymic, &student.EMail, &student.Password, &student.University, &student.Direction, &student.Group); err!=nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Student{}, errors.New("student not found")
		}
		return entity.Student{}, err
	}

	return entity.Student{}, nil
}