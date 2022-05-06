package server

import (
	"net/http"

	"github.com/ShiryaevNikolay/auth/internal/apperror"
	"github.com/ShiryaevNikolay/auth/internal/domain"
	"github.com/ShiryaevNikolay/auth/internal/res"
	"github.com/ShiryaevNikolay/auth/internal/utils"
)

func (server *Server) GetUserId(w http.ResponseWriter, r *http.Request) error {
	server.logger.Infoln(res.LogGetTokenFromHeader)
	token := r.Header.Get("Token")

	server.logger.Infoln(res.LogGetUserId)
	tokenModel, err := server.service.GetToken(token)
	if err != nil {
		server.logger.Errorf(res.LogCouldNotFindUserWithThisToken, err)
		return apperror.New(
			err, 
			res.ErrorAccessIsDenied, 
			http.StatusForbidden,
		)
	}

	server.logger.Infoln(res.LogConvertDataToJSON)
	err = utils.ResponseOk(w, domain.User{
		ID: tokenModel.UserID,
	})
	if err != nil {
		server.logger.Errorf(res.LogErrorConvertDataToJSON, err)
		return apperror.SystemError(err)
	}

	return nil
}
