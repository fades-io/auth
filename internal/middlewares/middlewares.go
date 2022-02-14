package middlewares

import (
	"errors"
	"net/http"

	"github.com/ShiryaevNikolay/auth/internal/apperror"
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
				if errors.Is(err, apperror.ErrNotFound) {
					w.WriteHeader(http.StatusNotFound)
					w.Write(apperror.ErrNotFound.Marshal())
					return
				}

				err = err.(*apperror.AppError)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(appErr.Marshal())
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
