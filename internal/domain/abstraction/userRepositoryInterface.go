package abstraction

import "schedule/internal/domain/models"

type UserRepositoryInterface interface {
	GetUserByUsername(username string) (*models.User, error)
}
