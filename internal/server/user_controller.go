package server

import (
	"net/http"

	"github.com/ShiryaevNikolay/auth/internal/apperror"
	"github.com/ShiryaevNikolay/auth/internal/domain"
	"github.com/ShiryaevNikolay/auth/internal/utils"
)

func (server *Server) GetUserId(w http.ResponseWriter, r *http.Request) error {
	server.logger.Infoln("Получение токена из Header")
	token := r.Header.Get("Token")

	server.logger.Infoln("Получение id пользователя")
	tokenModel, err := server.service.GetUserIdByToken(token)
	if err != nil {
		server.logger.Errorf("Не удалось найти пользователя с этим токеном", err)
		return apperror.New(err, "Доступ запрещен", err.Error(), http.StatusForbidden)
	}

	server.logger.Infoln("Конвертация данных в JSON")
	err = utils.ResponseOk(w, domain.User{
		ID: tokenModel.UserID,
	})
	if err != nil {
		server.logger.Errorf("Ошибка конвертации данных в JSON: %v", err)
		return apperror.SystemError(err)
	}

	return nil
}
