package res

const (
	ErrorSystem = "Внутренняя ошибка сервера"
	ErrorReadBody = "Не удалось считать тело запроса"
	ErrorConvertBodyToJSON = "Не удалось преобразовать JSON в модель"
	ErrorInvalidDataFormat =  "Неверный формат данных. Проверьте, корректно ли введен логин/пароль"
	ErrorUserNotFound =  "Пользователя с таким логином/паролем не существует"
	ErrorInvalidPassword =  "Неверный пароль"
	ErrorLoginRequired =  "требуется логин"
	ErrorPasswordRequired =  "требуется пароль"
	ErrorBadFormat =  "неверный формат"
)