package binCard

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Number struct {
	Length int  `json:"length"`
	Luhn   bool `json:"luhn"`
}

type Country struct {
	Numeric   string  `json:"numeric"`
	Alpha2    string  `json:"alpha2"`
	Name      string  `json:"name"`
	Emoji     string  `json:"emoji"`
	Currency  string  `json:"currency"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Bank struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Phone string `json:"phone"`
	City  string `json:"city"`
}

type BinFullInfo struct {
	Number  Number  `json:"number"`
	Scheme  string  `json:"scheme"`
	Type    string  `json:"type"`
	Brand   string  `json:"brand"`
	Prepaid bool    `json:"prepaid"`
	Country Country `json:"country"`
	Bank    Bank    `json:"bank"`
}

type BinInfo struct {
	Brand   string `json:"brand"`
	Type    string `json:"type"`
	Prepaid bool   `json:"prepaid"`
	BinBank string `json:"bin_bank"`
	Country string `json:"country"`
}

func checkFullBin(needBin string) (BinFullInfo, error) {
	response, err := http.Get(fmt.Sprintf("https://lookup.binlist.net/%s", needBin))
	if err != nil {
		return BinFullInfo{}, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return BinFullInfo{}, err
		}
		var binFullInfo BinFullInfo
		err = json.Unmarshal([]byte(body), &binFullInfo)
		if err != nil {
			fmt.Println("error unmarshal")
			return BinFullInfo{}, err
		}
		return binFullInfo, nil
	}
	return BinFullInfo{}, nil
}

func CheckBin(bin string) (BinInfo, error) {
	resCheckBin, err := checkFullBin(bin)
	if err != nil {
		return BinInfo{}, err
	}

	var binInfo BinInfo = BinInfo{
		Brand:   resCheckBin.Brand,
		Type:    resCheckBin.Type,
		Prepaid: resCheckBin.Prepaid,
		BinBank: resCheckBin.Bank.Phone,
		Country: resCheckBin.Country.Name,
	}

	return binInfo, nil
}
