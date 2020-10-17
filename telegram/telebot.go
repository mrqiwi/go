package main

import (
	"log"
    "os"
    "os/exec"
    "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

    bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
    if err != nil {
        log.Panic(err)
    }
	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

        log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

        switch update.Message.Text {
            case "off":
                if err := exec.Command("/sbin/shutdown", "-h", "now").Run(); err != nil {
                    log.Printf("Failed to initiate shutdown: %s", err)
                    msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Cannot execute this command")
                    bot.Send(msg)
                } else {
                    msg := tgbotapi.NewMessage(update.Message.Chat.ID, "OK")
                    bot.Send(msg)
                }

            case "reboot":
                if err := exec.Command("/sbin/shutdown", "-r", "now").Run(); err != nil {
                    log.Printf("Failed to initiate shutdown: %s", err)
                    msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Cannot execute this command")
                    bot.Send(msg)
                } else {
                    msg := tgbotapi.NewMessage(update.Message.Chat.ID, "OK")
                    bot.Send(msg)
                }

            default:
                msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I don't understand you")
                bot.Send(msg)
        }
    }
}
