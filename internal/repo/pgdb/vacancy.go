package pgdb

import (
	"context"

	"github.com/greenblat17/digital_spb/internal/entity"
	"github.com/greenblat17/digital_spb/pkg/postgres"

	log "github.com/sirupsen/logrus"
)

type VacancyRepo struct {
	*postgres.Postgres
}

func NewVacancyRepo(pg *postgres.Postgres) *VacancyRepo {
	return &VacancyRepo{pg}
}

func (r *VacancyRepo) CreateVacancy(ctx context.Context, vacancy entity.Vacancy) (int, error) {
	log.Info("Creating vacancy: ", vacancy)

	var id int
	query := "INSERT INTO vacancy (name) VALUES ($1) RETURNING id"

	row := r.Pool.QueryRow(ctx, query, vacancy.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *VacancyRepo) CountVacancy(ctx context.Context) (int, error) {
	query := "SELECT count(*) FROM vacancy"
	var count int
	row := r.Pool.QueryRow(ctx, query)
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (r *VacancyRepo) GetVacanciesByEducationId(ctx context.Context, educationId int) ([]entity.Vacancy, error) {
	query := "SELECT * FROM vacancy_education WHERE education_id = $1"
	rows, err := r.Pool.Query(ctx, query, educationId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vacEds []entity.VacancyEducation
	for rows.Next() {
		var vacEd entity.VacancyEducation
		if err := rows.Scan(&vacEd.Id, &vacEd.EducationId, &vacEd.VacancyId); err != nil {
			return nil, err
		}
		vacEds = append(vacEds, vacEd)
	}

	var vacancies []entity.Vacancy
	for _, vacEd := range vacEds {
		query = "SELECT * FROM vacancy WHERE vacancyId = $1"
		row := r.Pool.QueryRow(ctx, query, vacEd.VacancyId)

		var vacancy entity.Vacancy
		err := row.Scan(&vacancy.Id, &vacancy.Name)
		if err != nil {
			return nil, err
		}

		vacancies = append(vacancies, vacancy)
	}

	return vacancies, nil
}
