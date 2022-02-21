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
	server.logger.Infoln("Получение Body запроса")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		server.logger.Errorf("Ошибка получения Body: %v", err)
		return apperror.New(err, "Не удалось считать тело запроса", err.Error(), http.StatusUnprocessableEntity)
	}

	server.logger.Infoln("Конвертация Body в модель UserLogin")
	user := domain.UserLogin{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		server.logger.Errorf("Ошибка конвертации Body в UserModel: %v", err)
		return apperror.New(err, "Не удалось преобразовать JSON в модель", err.Error(), http.StatusUnprocessableEntity)
	}

	server.logger.Infoln("Подготовка и валидация пользовательских данных")
	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		server.logger.Errorf("Ошибка валидации пользовательских данных: %v", err)
		return apperror.New(err, "Неверный формат данных. Проверьте, корректно ли введен логин/пароль", err.Error(), http.StatusBadRequest)
	}

	server.logger.Infoln("Авторизация пользователя")
	token, err := server.SignIn(user.Username, user.Password)
	if err != nil {
		server.logger.Errorf("Ошибка авторизации пользователя: %v", err)
		return err
	}

	server.logger.Infoln("Конвертация данных в JSON")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		server.logger.Errorf("Ошибка конвертации данных в JSON: %v", err)
		return apperror.SystemError(err)
	}
	return nil
}

// Получает пользователя и создает токен
func (server *Server) SignIn(username, password string) (string, error) {
	server.logger.Infoln("Получение пользователя из БД")
	user, err := server.service.GetUser(username)
	if err != nil {
		server.logger.Errorf("Не удалось получить пользователя: %v", err)
		return "", apperror.New(err, "Пользователя с таким логином/паролем не существует", err.Error(), http.StatusNotFound)
	}

	server.logger.Infoln("Проверка пароля пользователя")
	err = domain.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		server.logger.Errorf("Пароли пользователя не совпадают: %v", err)
		return "", apperror.New(err, "Неверный пароль", err.Error(), http.StatusUnauthorized)
	}

	server.logger.Infoln("Создание токена для пользователя")
	token, err := auth.CreateToken(user.ID)
	if err != nil {
		server.logger.Errorf("Ошибка создания токена: %v", err)
		return "", apperror.SystemError(err)
	}

	tokenModel := domain.Token{
		Value:   token,
		Status:  "Created",
		UserID:  user.ID,
	}

	server.logger.Infoln("Создание токена в БД")
	err = server.service.CreateToken(&tokenModel)
	if err != nil {
		server.logger.Errorf("Ошибка создания токена в БД: %v", err)
		return "", apperror.SystemError(err)
	}

	return token, nil
}
