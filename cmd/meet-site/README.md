# http API

```sh
# dev
http://localhost:8080/
```

### Регистрация пользователя

````sh
Url: http://localhost:8080/auth/sign-up
```js
{
  "body": {
    "username": string, <- username пользователя
    "email": string, <- email пользователя
    "password": string <- password пользователя
  }
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "успешная регистрация пользователя" <- typeof string

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка конвертации userForm, {err}";
  -message: "ошибка выполнения функции uid из базы данных, {err}";
  -message: "ошибка выполнения функции user_create из базы данных, {err}";
````

### Авторизация пользователя

````sh
Url: http://localhost:8080/auth/sign-in
```js
{
  "body": {
    "username": string, <- username пользователя
    "password": string <- password пользователя
  }
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "успешная авторизация пользователя" <- typeof string
  -token: "..." <- typeof string

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка выполнения функции user_get_data из базы данных, {err}";
  -message: "ошибка конвертации в функции GetUser, {err}";
  -message: "пустой заголовок авторизации";
  -message: "неверный заголовок авторизации";
  -message: "идентификатор пользователя не найден";
  -message: "идентификатор пользователя имеет недопустимый тип";
  -message: "неправильный логин или пароль";
````

### Валидация электронной почты

````sh
Url: http://localhost:8080/validate/validateEmail
```js
{
  "body": {
    "email": string, <- email пользователя
  }
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "успешная валидация электронной почты" <- typeof string
  -result: true <- typeof bool

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "проверить адрес электронной почты не удалось, ошибка, {err}";
  -message: "синтаксис адреса электронной почты недействителен";
  -message: "ошибка валидации электронной почты";
````

### Валидация пароля

````sh
Url: http://localhost:8080/validate/validatePassword
```js
{
  "body": {
    "password": string, <- password пользователя
  }
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "успешная валидация пароля" <- typeof string
  -result: true <- typeof bool

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "пароль должен быть не менее 8 символов";
  -message: "пароль должен состоять как минимум из 2 цифр";
  -message: "пароль должен содержать не менее 4 буквенных символов";
  -message: "пароль должен содержать как минимум 1 специальный символ";
  -message: "пароль должен содержать как минимум 1 символ в верхнем регистре";
  -message: "пароль не должен содержать соседние символы с одинаковым значением";
  -message: "пароль содержит недопустимую комбинацию символов: 'asdf', 'qwerty', '1234' or '98765'";
  -message: "пароль содержит значения, следующие друг за другом, 1234, 3456, abcd, efgh";
  -message: "ошибка валидации пароля";
````

### Валидация username

````sh
Url: http://localhost:8080/validate/validateUsername
```js
{
  "body": {
    "username": string, <- username пользователя
  }
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "успешная валидация username" <- typeof string
  -result: true <- typeof bool

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "в имени пользователя разрешены только буквенно-цифровые символы";
  -message: "длина имени пользователя должна быть больше 4 и меньше 51 символа";
  -message: "ошибка валидации username";
````

### Верификация электронной почты

````sh
Url: http://localhost:8080/verify/emailver/:uid
```js
{
  "params": {
    "uid": string, <- uid пользователя
  }
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "успешное подтверждение электронной почты" <- typeof string
  -result: true <- typeof bool

POSSIBLE MISTAKES:
  -message: "ошибка выполнения функции user_verify_email из базы данных, {err}";
  -message: "ошибка подтверждения электронной почты";
````

### Отправка письма на почту пользователя для восстановления пароля

````sh
Url: http://localhost:8080/recovery/recoveryPasswordSendMessage
```js
{
  "body": {
    "email": string, <- email пользователя
  }
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "успешная отправка письма для восстановления пароля" <- typeof string
  -result: true <- typeof bool

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка конвертации userForm, {err}";
  -message: "ошибка выполнения функции user_get_uid_by_email из базы данных, {err}";
  -message: "ошибка отправки письма восстановления пароля на вашу почту";
  -message: "ошибка отправки письма для восстановления пароля";
````

### Изменение пароля пользователя

````sh
Url: http://localhost:8080/recovery/recoveryPassword
```js
{
  "body": {
    "uid": string, <- uid пользователя
    "password": string, <- пароль пользователя
  }
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "успешное изменение пароля" <- typeof string
  -result: true <- typeof bool

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка конвертации userForm, {err}";
  -message: "ошибка выполнения функции user_recovery_password из базы данных, {err}";
  -message: "ошибка изменения пароля";
````
