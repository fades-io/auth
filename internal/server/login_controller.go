package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ShiryaevNikolay/auth/internal/apperror"
	"github.com/ShiryaevNikolay/auth/internal/auth"
	"github.com/ShiryaevNikolay/auth/internal/domain"
	"github.com/ShiryaevNikolay/auth/internal/utils"
	"github.com/ShiryaevNikolay/auth/internal/res"
	"golang.org/x/crypto/bcrypt"
)

// Обработка запроса входа пользователя
func (server *Server) Login(w http.ResponseWriter, r *http.Request) error {
	server.logger.Infoln(res.LogGettingBodyRequest)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		server.logger.Errorf(res.LogErrorGettingBody, err)
		return apperror.New(
			err, 
			res.ErrorReadBody, 
			err.Error(), 
			http.StatusUnprocessableEntity,
		)
	}

	server.logger.Infoln(res.LogConvertBodyToUserLogin)
	user := domain.UserLogin{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		server.logger.Errorf(res.LogErrorConvertingBotyToUserLogin, err)
		return apperror.New(
			err, 
			res.ErrorConvertBodyToJSON, 
			err.Error(), 
			http.StatusUnprocessableEntity,
		)
	}

	server.logger.Infoln(res.LogPreparationAndValidationOfUserData)
	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		server.logger.Errorf(res.LogErrorValidationOfUserData, err)
		return err
	}

	server.logger.Infoln(res.LogUserAuthorization)
	token, err := server.SignIn(user.Username, user.Password)
	if err != nil {
		server.logger.Errorf(res.LogErrorUserAuthorization, err)
		return err
	}

	server.logger.Infoln(res.LogConvertDataToJSON)
	err = utils.ResponseOk(w, domain.Token{
		Value:   token,
	})
	if err != nil {
		server.logger.Errorf(res.LogErrorConvertDataToJSON, err)
		return apperror.SystemError(err)
	}
	return nil
}

// Получает пользователя и создает токен
func (server *Server) SignIn(username, password string) (string, error) {
	server.logger.Infoln(res.LogGettingUserFromDB)
	user, err := server.service.GetUser(username)
	if err != nil {
		server.logger.Errorf(res.LogErrorGetUser, err)
		return "", apperror.New(
			err, 
			res.ErrorUserNotFound, 
			err.Error(), 
			http.StatusNotFound,
		)
	}

	server.logger.Infoln(res.LogCheckUserPassword)
	err = domain.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		server.logger.Errorf(res.LogUserPasswordDoNotMatch, err)
		return "", apperror.New(
			err, 
			res.ErrorInvalidPassword, 
			err.Error(), 
			http.StatusUnauthorized,
		)
	}

	server.logger.Infoln(res.LogGenerateTokenForUser)
	token, err := auth.CreateToken(user.ID)
	if err != nil {
		server.logger.Errorf(res.LogErrorCreateToken, err)
		return "", apperror.SystemError(err)
	}

	tokenModel := domain.Token{
		Value:   token,
		Status:  domain.Created,
		UserID:  user.ID,
	}

	server.logger.Infoln(res.LogCreateTokenInDB)
	err = server.service.CreateToken(&tokenModel)
	if err != nil {
		server.logger.Errorf(res.LogErrorCreateTokenInDB, err)
		return "", apperror.SystemError(err)
	}

	err = server.SetDisabledAllTokens(user.ID, token)
	if err != nil {
		return "", nil
	}

	return token, nil
}

// Устанавливает всем токенам для данного пользователя статус "Disabled"
func (server *Server) SetDisabledAllTokens(userId uint, token string) error {
	err := server.service.UpdateStatusAllTokens(userId, token, domain.Disabled)
	if err != nil {
		return apperror.SystemError(err)
	}
	return nil
}
