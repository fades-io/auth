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
func (postgres *postgresDB) GetUser(username, password string) (*domain.User, error) {
	user := domain.User{}
	
	err := postgres.db.Debug().Table("users").Model(user).Where("username = ?", username).Take(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
