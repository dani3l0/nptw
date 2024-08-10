package telegram

import (
	tg "github.com/amarnathcjd/gogram/telegram"
)

func ReplyToUser(message *tg.NewMessage) {
	if !message.IsPrivate() {
		return
	}
	if message.Text() == "/start" {
		message.Reply("Czołem! Jestem botem, który wrzuca powtórki streamów NPTV na Telegrama i archive.org. Jeśli nie masz jak oglądać żywców, w opisie mojego profilu są kanały, gdzie wrzucam powtórki. Jakość może i dupy nie urywa, ale zawsze staram się znaleźć jak najlepszy format żeby zmieścić się w Telegramowym limicie.\n**💀 ŚMIERĆ WROGOM POLSKI! 💀**")
	}
}
