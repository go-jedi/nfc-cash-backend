package appl_row

type CreateUser struct {
	Username string `json:"username" binding:"required"`
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
}
