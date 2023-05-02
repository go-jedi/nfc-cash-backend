package appl_row

type User struct {
	Id            int    `json:"id"`
	Uid           string `json:"uid"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	IsVerifyEmail bool   `json:"is_verify_email"`
	Password      string `json:"password"`
	Role          string `json:"role"`
}

type UserProfile struct {
	Id       int    `json:"id"`
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
