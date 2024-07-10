package mocks

import (
	"schedule/internal/domain/dto"
	"schedule/internal/domain/models"
)

type MockClassMapper struct {
}

func (m *MockClassMapper) MapToDto(*models.Class) *dto.ClassDto {
	return &dto.ClassDto{}
}

func (m *MockClassMapper) MapToModel(*dto.ClassDto) *models.Class {
	return &models.Class{}
}

func (m *MockClassMapper) MapToCreateClassModel(*dto.CreateClassDto) *models.CreateClass {
	return &models.CreateClass{}
}

func (m *MockClassMapper) MapToUpdateClassModel(*dto.UpdateClassDto) *models.UpdateClass {
	return &models.UpdateClass{}
}
