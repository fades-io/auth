package domain

import "golang.org/x/crypto/bcrypt"

// Модель пользователя, которую получаем из БД
type User struct {
	ID       uint
	Username string
	Password string
}

// Проверка пароля пользователя
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
