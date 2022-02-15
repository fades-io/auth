package domain

// Модель пользователя, которую получаем из БД
type User struct {
	ID       uint
	Username string
	Password string
}
