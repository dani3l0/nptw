package telegram

import (
	"nptw/config"
	"nptw/utils"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func SendNotification(thumbnail interface{}, message string) {
	utils.Log("Sending notification about started stream to Telegram")
	_, err := client.SendMedia(config.Get().ChannelId, thumbnail, &tg.MediaOptions{
		Caption:   message,
		ParseMode: "markdown",
	})
	if err != nil {
		utils.Err("Couldn't send notification!")
		utils.Err(err.Error())
	}
}
