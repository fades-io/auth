package utils

import (
	"encoding/json"
	"net/http"

	"github.com/ShiryaevNikolay/auth/internal/apperror"
)

// Устанавливает статус 200 и возвращаем тело ответа
func ResponseOk(w http.ResponseWriter, body interface{}) error {
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(body)
	return err
}

// Устанавливает код статуса из ошибки и возвращает тело ошибки
func ResponseError(w http.ResponseWriter, body *apperror.AppError) {
	w.WriteHeader(body.Code)

	w.Write(body.Marshal())
}