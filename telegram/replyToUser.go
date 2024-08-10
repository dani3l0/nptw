package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"nptw/config"
	"nptw/utils"
	"strconv"
	"strings"

	tg "github.com/amarnathcjd/gogram/telegram"
)

var repliesSWO = []string{
	"💀 ŚMIERĆ WROGOM OJCZYZNY! 💀",
	"Śmierć! Śmierć! Śmierć!",
	"Śmierć skurwysynom!",
	"Śmierć kurwa!",
	"Śmierć!",
}

var repliesFight = []string{
	"Jeb go w ten pusty łeb!",
	"Szpadlem skurwysyna!",
	"Kijem go kurwa!",
	"W mordę kurwa!",
	"Bić!",
}

func ReplyToUser(message *tg.NewMessage) {
	sender := message.Sender
	if !message.IsPrivate() || sender.ID == myID {
		return
	}
	text := message.Text()
	chatId := message.ChatID()
	utils.Log(fmt.Sprint("Chat: received ", text, " from ", sender.ID, " (", sender.FirstName, ")"))

	if text == "/start" {
		keyboardMessage(
			chatId,
			"Czołem! Jestem botem, który wrzuca powtórki streamów NPTV na Telegrama i archive.org. Jeśli nie masz jak oglądać żywców, w opisie mojego profilu są kanały, gdzie wrzucam powtórki. Jakość może i dupy nie urywa, ale zawsze staram się znaleźć jak najlepszy format żeby zmieścić się w Telegramowym limicie.\n**💀 ŚMIERĆ WROGOM POLSKI! 💀**",
		)

	} else if text == "/stop" {
		keyboardMessage(chatId, "Już Ci nie zawracam dupy. Guziczki schowane. Bywaj!")

	} else {
		words := strings.Fields(text)
		if words[0] == "Bić" {
			randomIndex := rand.Intn(len(repliesFight))
			keyboardMessage(chatId, repliesFight[randomIndex])

		} else if words[0] == "Śmierć" {
			randomIndex := rand.Intn(len(repliesSWO))
			keyboardMessage(chatId, repliesSWO[randomIndex])
		}
	}
}

type TgMsgMarkup struct {
	ChatID      string      `json:"chat_id"`
	Text        string      `json:"text"`
	ParseMode   string      `json:"parse_mode"`
	ReplyMarkup interface{} `json:"reply_markup"`
}

func keyboardMessage(chatId int64, text string) (response *http.Response, err error) {
	var jsonData []byte
	data := TgMsgMarkup{
		ChatID: strconv.FormatInt(chatId, 10),
		Text:   text,
		ReplyMarkup: map[string]interface{}{
			"keyboard": [][]map[string]string{
				{
					{"text": "Śmierć Wrogom Ojczyzny!"},
					{"text": "Śmierć Wrogom Polski!"},
				}, {
					{"text": "Bić lewaka!"},
					{"text": "Bić posła!"},
				},
			},
		},
	}
	jsonData, _ = json.Marshal(data)
	response, err = http.Post(
		fmt.Sprint("https://api.telegram.org/bot", config.Get().TelegramBotToken, "/sendMessage"),
		"application/json", bytes.NewBuffer(jsonData),
	)
	return response, err
}
