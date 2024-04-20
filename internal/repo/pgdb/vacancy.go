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
	query := "INSERT INTO vacancy (name, education) VALUES ($1, $2) RETURNING id"

	row := r.Pool.QueryRow(ctx, query, vacancy.Name, vacancy.Education)
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
