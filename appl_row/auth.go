package appl_row

type CreateUser struct {
	Username string `json:"username" binding:"required"`
	TeleId   int64  `json:"tele_id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CheckEmailExist struct {
	Email string `json:"email"`
}

type CheckUsernameExist struct {
	Username string `json:"username"`
}

type CheckConfirmAccount struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
