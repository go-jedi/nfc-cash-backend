package appl_row

type Bot struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Token   string `json:"token"`
	ChatId  string `json:"chat_id"`
	IsAdmin bool   `json:"is_admin"`
	IsWork  bool   `json:"is_work"`
	Created string `json:"created"`
}

type BotCreate struct {
	Name   string `json:"name"`
	Token  string `json:"token"`
	ChatId string `json:"chat_id"`
}

type BotDelete struct {
	Token string `json:"token"`
}
