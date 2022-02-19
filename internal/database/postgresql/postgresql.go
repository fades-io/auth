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
	user := domain.User{}
	
	err := postgres.db.Debug().Table("users").Model(user).Where("username = ?", username).Take(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Создание токена
func (postgres *postgresDB) CreateToken(token *domain.Token) (error) {
	err := postgres.db.Debug().Table("tokens").Model(domain.Token{}).Create(token).Error
	if err != nil {
		return err
	}
	return nil
}
