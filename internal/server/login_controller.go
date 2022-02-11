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
func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Не удалось считать тело запроса: %v", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	user := domain.UserLogin{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatalf("Не удалось смапить json в domain модель: %v", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		log.Fatalf("Не удалось смапить json в domain модель: %v", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	token, err := server.SignIn(user.Username, user.Password)
	if err != nil {
		log.Fatalf("Не удалось смапить json в domain модель: %v", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func (server *Server) SignIn(username, password string) (string, error) {
	// TODO: получать пользователя из БД и проверять пароли
	// var err error
	// user := domain.UserLogin{}
	return auth.CreateToken(0)
}
