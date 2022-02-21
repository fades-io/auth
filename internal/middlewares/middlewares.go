package middlewares

import (
	"errors"
	"net/http"

	"github.com/ShiryaevNikolay/auth/internal/apperror"
	"github.com/ShiryaevNikolay/auth/internal/utils"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

// Добавляет заголовки запросу
func SetHeadersMiddleware(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var appErr *apperror.AppError
		err := h(w, r)
		if err != nil {
			/*
				Смотрим, ошибка наша, т.е. AppError или какая-то другая
			*/
			if errors.As(err, &appErr) {
				appErr := err.(*apperror.AppError)
				utils.ResponseError(w, appErr)
				return
			}

			w.WriteHeader(http.StatusTeapot)
			/*
				Таким образом получаем все системные ошибки обернутые в наш AppError
			*/
			w.Write(apperror.SystemError(err).Marshal())
		}
	}
}
