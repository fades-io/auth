package server

import (
	"github.com/ShiryaevNikolay/auth/internal/domain"
)

// Методы, которые должны реализовывать БД
type Storage interface {
	GetUser(username string) (*domain.User, error)
	CreateToken(token *domain.Token) (error)
	UpdateStatusAllTokens(userId uint, token, status string) error
}
