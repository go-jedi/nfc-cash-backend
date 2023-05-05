package appl_row

type GetUsersUnConfirm struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	TeleId   int64  `json:"tele_id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type GetUsersConfirm struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	TeleId   int64  `json:"tele_id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type UserConfirmAccount struct {
	Id int `json:"id"`
}

type ChangeUser struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	TeleId   int64  `json:"tele_id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
