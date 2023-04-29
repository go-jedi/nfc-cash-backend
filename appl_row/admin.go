package appl_row

type GetUsersUnConfirm struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type UserConfirmAccount struct {
	Id int `json:"id"`
}
