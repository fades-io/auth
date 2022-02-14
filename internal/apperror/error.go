package apperror

import "encoding/json"

var (
	ErrNotFound = New(nil, "Not found", "", "404")
)

// Кастомная ошибка, которая передается в json вместе с сообщением
type AppError struct {
	Err              error  `json:"-"` // исходная ошибка, поэтому в JSON она нам не нужна
	Message          string `json:"message,omitempty"`
	DeveloperMEssage string `json:"developer_message,omitempty"`
	Status           string `json:"code,omitempty"`
}

// метод для соответствия интерфейсу Error{}
func (appError *AppError) Error() string {
	return appError.Message
}

func (appError *AppError) Unwrap() error {
	return appError.Err
}

func (appError *AppError) Marshal() []byte {
	marshal, err := json.Marshal(appError)
	if err != nil {
		return nil
	}
	return marshal
}

// Создает кастомную ошибку
func New(err error, message, developerMessage, code string) *AppError {
	return &AppError{
		Err:              err,
		Message:          message,
		DeveloperMEssage: developerMessage,
		Status:           code,
	}
}

// Любую ошибку оборачиваем в системную
func SystemError(err error) *AppError {
	return New(err, "internal system error", err.Error(), "")
}
