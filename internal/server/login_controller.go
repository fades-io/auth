package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ShiryaevNikolay/auth/internal/apperror"
	"github.com/ShiryaevNikolay/auth/internal/auth"
	"github.com/ShiryaevNikolay/auth/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

// Обработка запрос для входа пользователя
func (server *Server) Login(w http.ResponseWriter, r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return apperror.New(err, "Не удалось считать тело запроса", err.Error(), http.StatusUnprocessableEntity)
	}

	user := domain.UserLogin{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return apperror.New(err, "Не удалось преобразовать JSON в модель", err.Error(), http.StatusUnprocessableEntity)
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		return apperror.New(err, "Неверный формат данных. Проверьте, корректно ли введен логин/пароль", err.Error(), http.StatusBadRequest)
	}
	token, err := server.SignIn(user.Username, user.Password)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		return apperror.SystemError(err)
	}
	return nil
}

// Получает пользователя и создает токен
func (server *Server) SignIn(username, password string) (string, error) {
	user, err := server.service.GetUser(username)
	if err != nil {
		return "", apperror.New(err, "Пользователя с таким логином/паролем не существует", err.Error(), http.StatusNotFound)
	}

	err = domain.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", apperror.New(err, "Неверный пароль", err.Error(), http.StatusUnauthorized)
	}

	token, err := auth.CreateToken(user.ID)
	if err != nil {
		return "", apperror.SystemError(err)
	}

	tokenModel := domain.Token{
		Value:   token,
		Status:  "Created",
		UserID:  user.ID,
	}

	err = server.service.CreateToken(&tokenModel)
	if err != nil {
		return "", apperror.SystemError(err)
	}

	return token, nil
}
