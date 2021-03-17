package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	Handler()
}

func Handler() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// msg := tgbotapi.NewMessage(-582130977, "тест")
	//
	// bot.Send(msg)

	// info, err := bot.GetWebhookInfo()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if info.LastErrorDate != 0 {
	// 	log.Printf("[Telegram callback failed]%s", info.LastErrorMessage)
	// }

	http.HandleFunc("/"+bot.Token, func(w http.ResponseWriter, r *http.Request) {
		update, _ := bot.HandleUpdate(r)

		log.Printf("%+v\n", *update)

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "status":
				msg.Text = "I'm ok."
			default:
				msg.Text = "I don't know that command"
			}
			bot.Send(msg)
		}
	})
}
