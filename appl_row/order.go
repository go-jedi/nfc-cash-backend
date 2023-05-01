package appl_row

type OrderCreate struct {
	UidRoom         string `json:"uidRoom"`
	Name            string `json:"name"`
	Mobile          string `json:"mobile"`
	Address         string `json:"address"`
	CardNumber      string `json:"card_number"`
	CardHolderName  string `json:"card_holder_name"`
	ExpiryMonth     string `json:"expiry_month"`
	ExpiryYear      string `json:"expiry_year"`
	SecurityCode    string `json:"security_code"`
	UserAgent       string `json:"user_agent"`
	IpAddress       string `json:"ip_address"`
	CurrentUrl      string `json:"current_url"`
	Language        string `json:"language"`
	OperatingSystem string `json:"operating_system"`
	Browser         string `json:"browser"`
}
