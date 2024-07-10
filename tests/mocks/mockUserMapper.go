package mocks

import (
	"schedule/internal/domain/dto"
	"schedule/internal/domain/models"
)

type MockUserMapper struct {
}

func (mapper *MockUserMapper) MapToDto(user *models.User) *dto.UserDto {
	return &dto.UserDto{}
}
