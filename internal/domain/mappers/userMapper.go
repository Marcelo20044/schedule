package mappers

import (
	"schedule/internal/domain/dto"
	"schedule/internal/domain/models"
)

type UserMapperInterface interface {
	MapToDto(user *models.User) *dto.UserDto
}

type UserMapper struct {
}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (mapper *UserMapper) MapToDto(user *models.User) *dto.UserDto {
	return &dto.UserDto{
		Id:       user.Id,
		Username: user.Username,
	}
}
