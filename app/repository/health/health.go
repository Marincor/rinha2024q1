package health

import (
	"time"

	"api.default.marincor.com/adapters/database"
	"api.default.marincor.com/entity"
)

type Repository struct{}

func New() *Repository {
	return &Repository{}
}

func (repo *Repository) Insert(now time.Time) error {
	return database.Exec("INSERT INTO health (sync) VALUES ($1)", now)
}

func (repo *Repository) GetOne(now time.Time) (*entity.Health, error) {
	return database.New[*entity.Health]().QueryOne("SELECT * FROM health WHERE sync = $1", new(entity.Health), now)
}
