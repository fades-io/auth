package domain

import (
	"errors"
	"html"
	"regexp"
	"strings"

	"github.com/ShiryaevNikolay/auth/internal/res"
)

// Данные пользователя для авторизации
type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Подготовка информации о пользователе
func (user *UserLogin) Prepare() {
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	user.Password = html.EscapeString(strings.TrimSpace(user.Password))
}

// Валидация информации о пользователе
func (user *UserLogin) Validate(action string) error {
	switch strings.ToLower(action) {
	case "login":
		if user.Username == "" {
			return errors.New(res.ErrorLoginRequired)
		}
		if user.Password == "" {
			return errors.New(res.ErrorPasswordRequired)
		}
		return nil
	default:

		return nil
	}
}

var (
	ErrBadFormat = errors.New(res.ErrorBadFormat)
	emailRegexp  = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

// Валидация почты
func ValidateFormat(email string) error {
	if !emailRegexp.MatchString(email) {
		return ErrBadFormat
	}
	return nil
}
