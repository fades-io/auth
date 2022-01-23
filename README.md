# Микросервис Авторизация

## POST /user/login

Запрос POST, так как при авторизации создается токен для пользователя

Тело запроса
``` json
{
    "username": "string",
    "password": "string"
}
```
Ответ: пользователь найден, код **200**
``` json
{
    "user_token": "string"
}
```
Если пользователь найден, возвращается его токен.

Ответ: пользователь **не** найден, код **200**
``` json
{
    "error": {
        "code": 401,
        "message": "Такой пользователь не существует. Проверьте, корректно ли введены данные"
    }
}
```

## GET /user/token

Параметры запроса
```json
{
    "token": "example_token"
}
```

Ответ: информация о токене, код **200**
``` json
{
    "status": string,
    "created": string,
    "expired": 3600
}
```
- `status` - статус токена
- `created` - дата создания
- `expired` - длительность жизни токена

## DELETE /user/token

Параметры запроса
``` json
{
    "token": "example_tokenудаляет токен"
}
```

Ответ: код **204**

Данный запрос делает токен недействительным