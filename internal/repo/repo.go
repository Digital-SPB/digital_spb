package repo

import "github.com/greenblat17/digital_spb/pkg/postgres"

type Repositories struct {
}

func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{}
}
