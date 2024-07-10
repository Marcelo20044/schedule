package services

import (
	"schedule/internal/domain/abstraction"
	"schedule/internal/domain/dto"
	"schedule/internal/domain/mappers"
)

type UserService struct {
	Repository abstraction.UserRepositoryInterface
	Mapper     mappers.UserMapperInterface
}

func NewUserService(repository abstraction.UserRepositoryInterface, mapper mappers.UserMapperInterface) *UserService {
	return &UserService{Repository: repository, Mapper: mapper}
}

func (service *UserService) GetUserByUsername(username string) (*dto.UserDto, error) {
	user, err := service.Repository.GetUserByUsername(username)
	return service.Mapper.MapToDto(user), err
}
