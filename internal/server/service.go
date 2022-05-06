package server

import (
	"github.com/ShiryaevNikolay/auth/internal/domain"
)

// Сущность, которая содержит в себе интерфейс для работы с БД
type service struct {
	storage Storage
}

// Сервис для работы с БД
type Service interface {
	GetUser(username string) (*domain.User, error)
	CreateToken(token *domain.Token) (error)
	UpdateStatusAllTokens(userId uint, token, status string) error
}

// Конструктор для создания сервиса
func NewService(storage Storage) Service {
	return &service{
		storage: storage,
	}
}

// Получение пользователя по имени
func (s *service) GetUser(username string) (*domain.User, error) {
	return s.storage.GetUser(username)
}

// Создание токена в БД
func (s *service) CreateToken(token *domain.Token) (error) {
	return s.storage.CreateToken(token)
}

// Обновление статуса у всех токенов
func (s *service) UpdateStatusAllTokens(userId uint, token, status string) error {
	return s.storage.UpdateStatusAllTokens(userId, token, status)
}
