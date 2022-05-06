package postgresql

import (
	"github.com/ShiryaevNikolay/auth/internal/domain"
	"github.com/ShiryaevNikolay/auth/internal/server"
	"gorm.io/gorm"
)

const (
	tokensTable = "tokens"
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

	err := postgres.db.Debug().Table(
		"users",
	).Model(
		user,
	).Where(
		"username = ?",
		username,
	).Take(
		&user,
	).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Создание токена
func (postgres *postgresDB) CreateToken(token *domain.Token) error {
	err := postgres.db.Debug().Table(
		tokensTable,
	).Model(
		domain.Token{},
	).Create(
		token,
	).Error
	if err != nil {
		return err
	}
	return nil
}

// Обновление статуса у всех токенов
func (postgres *postgresDB) UpdateStatusAllTokens(userId uint, token, status string) error {
	err := postgres.db.Debug().Table(
		tokensTable,
	).Model(
		domain.Token{},
	).Where(
		"user_id = ?",
		userId,
	).Where(
		"token <> ?",
		token,
	).Update(
		"token_status",
		status,
	).Error
	if err != nil {
		return err
	}
	return nil
}

// Получение id пользователя по токену
func (postgres *postgresDB) GetToken(token string) (*domain.Token, error) {
	tokenModel := domain.Token{}
	err := postgres.db.Debug().Table(
		tokensTable,
	).Model(
		tokenModel,
	).Where(
		"token = ?",
		token,
	).Where(
		"token_status = 'Created'",
	).Take(
		&tokenModel,
	).Error
	if err != nil {
		return nil, err
	}
	return &tokenModel, nil
}
