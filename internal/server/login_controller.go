package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ShiryaevNikolay/auth/internal/auth"
	"github.com/ShiryaevNikolay/auth/internal/domain"
)

// Обработка запрос для входа пользователя
func (server *Server) Login(w http.ResponseWriter, r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Не удалось считать тело запроса: %v", err)
		return err
	}

	user := domain.UserLogin{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatalf("Не удалось смапить json в domain модель: %v", err)
		return err
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		log.Fatalf("Не удалось провалидировать поля: %v", err)
		return err
	}
	token, err := server.SignIn(user.Username, user.Password)
	if err != nil {
		log.Fatalf("Не удалось создать токен: %v", err)
		return err
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
		return err
	}
	return nil
}

func (server *Server) SignIn(username, password string) (string, error) {
	user, err := server.service.GetUser(username, password)
	if err != nil {
		return "", err
	}
	return auth.CreateToken(user.ID)
}
