package appl_row

type RecoveryPassword struct {
	Uid      string `json:"uid"`
	Password string `json:"password"`
}

type RecoveryPasswordSendMessage struct {
	Email string `json:"email"`
}

type CheckRecoveryPassword struct {
	Uid string `json:"uid"`
}

type CompleteRecoveryPassword struct {
	Uid string `json:"uid"`
}

type RecoveryPasswordCompare struct {
	Uid      string `json:"uid"`
	Password string `json:"password"`
}
