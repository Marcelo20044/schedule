package repositories

import (
	"github.com/jmoiron/sqlx"
	"schedule/internal/domain/models"
)

type UserRepository struct {
	Db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (repository *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	query := `
	select id, username, password
	from schedule.users
	where username = $1
	`
	user := new(models.User)
	err := repository.Db.Get(user, query, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}
