package functions

import (
	"fmt"
	"nptw/config"
	"nptw/providers/telegram"
	"strings"
)

func SendNotification(thumbnail string, title string) {
	telegram.SendNotification(thumbnail, fmt.Sprintf(
		"🇵🇱 Rozpoczął się Żywiec! 🇵🇱\n\n🐺 **%s** 🦎\n\n🔗 Link do DLive: https://dlive.tv/%s",
		strings.Replace(fmt.Sprintf("%v", title), "\n", "", -1), config.Get().DliveUsername,
	))
}
