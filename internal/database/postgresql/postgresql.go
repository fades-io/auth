package postgresql

import (
	"github.com/ShiryaevNikolay/auth/internal/domain"
	"github.com/ShiryaevNikolay/auth/internal/server"
	"gorm.io/gorm"
)

// Обертка над gorm.DB
type postgresDB struct {
	db *gorm.DB
}

// Конструктор БД
func New(db *gorm.DB) server.Storage {
	return &postgresDB{
		db: db,
	}
}

// Получение пользователя по имени
func (postgres *postgresDB) GetUser(username string) (*domain.User, error) {
	// TODO: обращаться к БД и получать пользователя
	return nil, nil
}
