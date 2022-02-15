package server

import "github.com/ShiryaevNikolay/auth/internal/domain"

// Сущность, которая содержит в себе интерфейс для работы с БД
type service struct {
	storage Storage
}

// Сервис для работы с БД
type Service interface {
	GetUser(username, password string) (*domain.User, error)
}

// Конструктор для создания сервиса
func NewService(storage Storage) Service {
	return &service{
		storage: storage,
	}
}

// Получение пользователя по имени
func (s *service) GetUser(username, password string) (*domain.User, error) {
	return s.storage.GetUser(username, password)
}
