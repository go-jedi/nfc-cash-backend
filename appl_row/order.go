package appl_row

type Orders struct {
	Id       int    `json:"id"`
	UidOrder string `json:"uid_order"`
	Status   string `json:"status"`
	Created  string `json:"created"`
}

type Order struct {
	Id              int    `json:"id"`
	UidOrder        string `json:"uid_order"`
	Status          string `json:"status"`
	BinBrand        string `json:"bin_brand"`
	BinType         string `json:"bin_type"`
	BinBank         string `json:"bin_bank"`
	BinCountry      string `json:"bin_country"`
	Name            string `json:"name"`
	Mobile          string `json:"mobile"`
	Address         string `json:"address"`
	CardNumber      string `json:"card_number"`
	CardHolderName  string `json:"card_holder_name"`
	ExpiryMonth     string `json:"exp_month"`
	ExpiryYear      string `json:"exp_year"`
	SecurityCode    string `json:"security_code"`
	UserAgent       string `json:"user_agent"`
	IpAddress       string `json:"ip_address"`
	CurrentUrl      string `json:"current_url"`
	Language        string `json:"lang"`
	OperatingSystem string `json:"operating_system"`
	Browser         string `json:"browser"`
	Created         string `json:"created"`
}

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

type GetOrder struct {
	UidOrder string `json:"uid_order"`
}
