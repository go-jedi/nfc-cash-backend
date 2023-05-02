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
    "tele_id": int64, <- telegram id пользователя
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

### Проверка существование email в базе данных

````sh
Url: http://localhost:8080/auth/check-email-exist
```js
{
  "body": {
    "email": string, <- email пользователя
  }
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "пользователь с такой электронной почтой уже существует" <- typeof string
  -result: true <- typeof string
OR:
RETURN:
  -status: 200 <- typeof int
  -message: "пользователь с такой электронной почтой не существует" <- typeof string
  -result: false <- typeof string

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка конвертации userForm, {err}";
  -message: "ошибка выполнения функции user_check_exist_email из базы данных, {err}";
````

### Проверка существование username в базе данных

````sh
Url: http://localhost:8080/auth/check-username-exist
```js
{
  "body": {
    "username": string, <- username пользователя
  }
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "пользователь с таким username уже существует" <- typeof string
  -result: true <- typeof string
OR:
RETURN:
  -status: 200 <- typeof int
  -message: "пользователь с таким username не существует" <- typeof string
  -result: false <- typeof string

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка конвертации userForm, {err}";
  -message: "ошибка выполнения функции user_check_exist_username из базы данных, {err}";
````

### Проверка подтверждение аккаунта пользователя администратором

````sh
Url: http://localhost:8080/auth/check-confirm-account
```js
{
  "body": {
    "username": string, <- username пользователя
    "password": string, <- пароль пользователя
  }
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "аккаунт пользователя успешно подтверждён администратором" <- typeof string
  -result: true <- typeof string
OR:
RETURN:
  -status: 200 <- typeof int
  -message: "аккаунт пользователя ещё не подтверждён администратором" <- typeof string
  -result: false <- typeof string

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка конвертации userForm, {err}";
  -message: "ошибка выполнения функции user_check_confirm_account из базы данных, {err}";
````

### Валидация электронной почты

````sh
Url: http://localhost:8080/validate/validate-email
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
Url: http://localhost:8080/validate/validate-password
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
Url: http://localhost:8080/validate/validate-username
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

### Проверка на верификацию электронной почты

````sh
Url: http://localhost:8080/verify/checkEmailVerify
```js
{
  "body": {
    "uid": string, <- uid пользователя
  }
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "электронная почта подтверждена" <- typeof string
  -result: true <- typeof bool
OR:
  -status: 200 <- typeof int
  -message: "электронная почта не подтверждена" <- typeof string
  -result: false <- typeof bool

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка выполнения функции user_verify_email из базы данных, {err}";
  -message: "ошибка подтверждения электронной почты";
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
Url: http://localhost:8080/recovery/recovery-password-send-message
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

### Сравнение паролей пользователя

````sh
Url: http://localhost:8080/recovery/recovery-password-compare
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
  -message: "успешное сравнение паролей" <- typeof string
  -result: true <- typeof bool
OR:
  -status: 200 <- typeof int
  -message: "ошибка сравнения паролей" <- typeof string
  -result: false <- typeof bool

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка выполнения функции user_compare_password из базы данных, {err}";
  -message: "ошибка сравнения паролей";
````

### Проверка запуска восстановления пароля

````sh
Url: http://localhost:8080/recovery/check-recovery-password
```js
{
  "body": {
    "uid": string, <- uid пользователя
  }
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "запуск восстановления пароля запущен" <- typeof string
  -result: true <- typeof bool
OR:
  -status: 200 <- typeof int
  -message: "запуск восстановления пароля не запущен" <- typeof string
  -result: false <- typeof bool

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка выполнения функции user_check_recovery_password из базы данных, {err}";
  -message: "запуск восстановления пароля не запущен";
````

### Завершение восстановления пароля

````sh
Url: http://localhost:8080/recovery/recovery-password-complete
```js
{
  "body": {
    "uid": string, <- uid пользователя
  }
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "успешное завершение восстановления пароля" <- typeof string

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка выполнения функции user_complete_recovery_password из базы данных, {err}";
````

### Изменение пароля пользователя

````sh
Url: http://localhost:8080/recovery/recovery-password
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

### Получение профиля пользователя

````sh
Url: http://localhost:8080/user/get-user-profile
```js
{
  "body": {}
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "успешное получение профиля пользователя" <- typeof string
  -result: []appl_row.UserProfile{} <- typeof bool

POSSIBLE MISTAKES:
  -message: "пустой заголовок авторизации";
  -message: "неверный заголовок авторизации";
  -message: "идентификатор пользователя не найден";
  -message: "идентификатор пользователя имеет недопустимый тип";
  -message: "ошибка выполнения функции user_get_profile из базы данных, {err}";
  -message: "ошибка конвертации в функции GetUserProfile";
````

### Проверка пользователя на администратора

````sh
Url: http://localhost:8080/user/check-is-admin
```js
{
  "body": {}
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "пользователь является администратором" <- typeof string
  -result: true <- typeof bool
OR:
  -status: 200 <- typeof int
  -message: "пользователь не является администратором" <- typeof string
  -result: false <- typeof bool

POSSIBLE MISTAKES:
  -message: "пустой заголовок авторизации";
  -message: "неверный заголовок авторизации";
  -message: "идентификатор пользователя не найден";
  -message: "идентификатор пользователя имеет недопустимый тип";
  -message: "ошибка выполнения функции user_check_is_admin из базы данных, {err}";
````

### Получение пользователей с подтвержденными аккаунтами (без супер администратора)

````sh
Url: http://localhost:8080/admin/get-users-confirm
```js
{
  "body": {}
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "успешное получение подтвержденных пользователей" <- typeof string
  -result: []appl_row.GetUsersConfirm{} <- typeof bool
OR:
  -status: 200 <- typeof int
  -message: "успешное получение подтвержденных пользователей" <- typeof string
  -result: []appl_row.GetUsersConfirm{} <- typeof bool

POSSIBLE MISTAKES:
  -message: "пустой заголовок авторизации";
  -message: "неверный заголовок авторизации";
  -message: "идентификатор пользователя не найден";
  -message: "идентификатор пользователя имеет недопустимый тип";
  -message: "ошибка выполнения функции admin_get_users_confirm из базы данных, {err}";
  -message: "ошибка конвертации в функции GetUsersConfirm, {err}";
````

### Получение пользователей с не подтвержденными аккаунтами

````sh
Url: http://localhost:8080/admin/get-users-un-confirm
```js
{
  "body": {}
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "успешное получение не подтвержденных пользователей" <- typeof string
  -result: []appl_row.GetUsersUnConfirm{} <- typeof bool
OR:
  -status: 200 <- typeof int
  -message: "успешное получение не подтвержденных пользователей" <- typeof string
  -result: []appl_row.GetUsersUnConfirm{} <- typeof bool

POSSIBLE MISTAKES:
  -message: "пустой заголовок авторизации";
  -message: "неверный заголовок авторизации";
  -message: "идентификатор пользователя не найден";
  -message: "идентификатор пользователя имеет недопустимый тип";
  -message: "ошибка выполнения функции admin_get_users_unconfirm из базы данных, {err}";
  -message: "ошибка конвертации в функции GetUsersUnConfirm, {err}";
````

### Подтверждение аккаунта пользователя администратором

````sh
Url: http://localhost:8080/admin/user-confirm-account
```js
{
  "body": {
	"id": int, <- id пользователя
  }
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "успешное подтверждение аккаунта пользователя" <- typeof string
  -result: true <- typeof bool
OR:
  -status: 200 <- typeof int
  -message: "ошибка подтверждения аккаунта пользователя" <- typeof string
  -result: false <- typeof bool

POSSIBLE MISTAKES:
  -message: "пустой заголовок авторизации";
  -message: "неверный заголовок авторизации";
  -message: "идентификатор пользователя не найден";
  -message: "идентификатор пользователя имеет недопустимый тип";
  -message: "некорректно переданы данные в body";
  -message: "ошибка конвертации userForm, {err}";
  -message: "ошибка выполнения функции admin_user_confirm_account из базы данных, {err}";
````

### Создание комнаты для чата

````sh
Url: http://localhost:8080/room/create-room
```js
{
  none
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "успешное создание комнаты" <- typeof string
  -result: true <- typeof bool
OR:
  -status: 200 <- typeof int
  -message: "ошибка создания комнаты" <- typeof string
  -result: false <- typeof bool

POSSIBLE MISTAKES:
  -message: "ошибка выполнения функции room_uid из базы данных, {err}";
  -message: "ошибка выполнения функции room_create из базы данных, {err}";
````

### Вступить в нужную комнату (чат)

````sh
Url: ws://localhost:8080/room/join-room/:roomId?uidUser=
```js
{
	"params": {
		"roomId": string, <- uid комнаты
	},
  "query": {
	"uidUser": string, <- uid пользователя (по умолчанию 'none', если у нас нету uid)
  }
}
```sh
RETURN:
	подключение к websockets

POSSIBLE MISTAKES:
  -message: "ошибка выполнения функции room_user_uid из базы данных, {err}";
  -message: "ошибка выполнения функции room_join из базы данных, {err}";
````

### Покинуть комнату (чат) пользователем

````sh
Url: http://localhost:8080/room/leave-room
```js
{
  "body": {
    "uidRoom": string, <- uid комнаты
    "uidUser": string, <- uid пользователя
  }
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "пользователь успешно покинул чат" <- typeof string

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка выполнения функции room_leave из базы данных, {err}";
````

### Покинуть комнату (чат) пользователем

````sh
Url: http://localhost:8080/room/get-room
```js
{
  "body": {
    "uidRoom": string, <- uid комнаты
  }
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "успешное получение комнаты" <- typeof string
  -result: []appl_row.Room{} <- typeof []appl_row.Room
OR:
  -status: 200 <- typeof int
  -message: "успешное получение комнаты" <- typeof string
  -result: []appl_row.Room{} <- typeof []appl_row.Room

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка выполнения функции room_get из базы данных, {err}";
  -message: "ошибка конвертации в функции GetRoom, {err}";
````

### Создание сообщения в чате

````sh
Url: http://localhost:8080/message/create-message
```js
{
  "body": {
    "uidRoom": string, <- uid комнаты
    "uidUser": string, <- uid пользователя
    "message": string, <- сообщение
  }
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "успешное создание сообщения" <- typeof string
  -result: true <- typeof bool
OR:
  -status: 200 <- typeof int
  -message: "ошибка создания сообщения" <- typeof string
  -result: false <- typeof bool

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка конвертации messageForm, {err}";
  -message: "ошибка выполнения функции message_create из базы данных, {err}";
````

### Получение всех сообщений нужного чата

````sh
Url: http://localhost:8080/message/get-room-messages
```js
{
  "body": {
    "uidRoom": string, <- uid комнаты
  }
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "успешное получение сообщений чата" <- typeof string
  -result: []appl_row.GetRoomMessages{} <- typeof bool
OR:
  -status: 200 <- typeof int
  -message: "успешное получение сообщений чата" <- typeof string
  -result: []appl_row.GetRoomMessages{} <- typeof bool

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка конвертации messageForm, {err}";
  -message: "ошибка выполнения функции messages_get_room из базы данных, {err}";
  -message: "ошибка конвертации в функции GetRoomMessages, {err}";
````

### Создание ордера (заказа)

````sh
Url: http://localhost:8080/order/create-order
```js
{
  "body": {
    "uidRoom": string, <- uid комнаты
    "name": string, <- имя
    "mobile": string, <- телефон
    "address": string, <- адрес
    "card_number": string, <- номер карты
    "card_holder_name": string, <- имя кому привязана карта
    "expiry_month": string, <- месяц просрочки карты
    "expiry_year": string, <- год просрочки карты
    "security_code": string, <- секретные 3 числа карты
    "user_agent": string, <- user_agent браузера
    "ip_address": string, <- ip адрес пользователя
    "current_url": string, <- текущий url адрес
    "language": string, <- язык браузера
    "operating_system": string, <- операционная система браузера
    "browser": string, <- название браузера
  }
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "успешное создание заказа" <- typeof string
  -result: true <- typeof bool
OR:
  -status: 200 <- typeof int
  -message: "ошибка создания заказа" <- typeof string
  -result: false <- typeof bool

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка выполнения функции CheckBin, {err}";
  -message: "ошибка конвертации orderForm, {err}";
  -message: "ошибка конвертации resCheckBin, {err}";
  -message: "ошибка выполнения функции order_create из базы данных, {err}";
````

### Получить нужный ордер (заказ)

````sh
Url: http://localhost:8080/order/get-order
```js
{
  "body": {
    "uid_order": string, <- uid заказа
  }
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "успешное получение заказа" <- typeof string
  -result: []appl_row.Order{} <- typeof bool
OR:
  -status: 200 <- typeof int
  -message: "успешное получение заказа" <- typeof string
  -result: []appl_row.Order{} <- typeof bool

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка выполнения функции order_get из базы данных, {err}";
  -message: "ошибка конвертации в функции GetUsersConfirm, {err}";
````

### Получить нужный ордер (заказ)

````sh
Url: http://localhost:8080/order/get-orders
```js
{
  "body": {}
}
```sh
RETURN:
  -status: 200 <- typeof int
  -message: "успешное получение заказов" <- typeof string
  -result: []appl_row.Order{} <- typeof bool
OR:
  -status: 200 <- typeof int
  -message: "успешное получение заказов" <- typeof string
  -result: []appl_row.Order{} <- typeof bool

POSSIBLE MISTAKES:
  -message: "ошибка выполнения функции orders_get из базы данных, {err}";
  -message: "ошибка конвертации в функции GetOrders, {err}";
````
