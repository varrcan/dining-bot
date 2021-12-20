package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/varrcan/workcalendar/europe/russia"
)

func main() {
	Handler()
}

func Handler() (string, error) {
	now := time.Now()
	weekday := now.Weekday()
	isFirstRun := now.Format("15:04") == "12:00"

	fmt.Println(isFirstRun)

	promp := []string{
		"В половину го",
		"Выступаем в 12:30",
		"Готовность полчаса",
		"Идем в половину",
		"Го в половину",
		"Ну что, го в половину",
		"Выдвигаемся в половину",
		"В половину обед",
		"В 12:30 на обед",
		"Обед в 12:30 или че?",
		"12:30 - обед!",
	}

	run := []string{
		"го!",
		"Go!",
		"Гоу гоу гоу",
		"Поiхали!",
		"Жратикоо",
		"На обед!",
		"Идём?",
		"Пора обедать",
		"Ну что, гоу?",
		"Время хавать, пошли",
		"Кушать го",
		"Го в столовку",
	}

	if (int(weekday) != 0 && int(weekday) != 6) && !russia.IsHoliday(now) {
		bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
		if err != nil {
			log.Panic(err)
		}

		var text string
		if isFirstRun {
			text = getRandom(promp)
		} else {
			text = getRandom(run)
		}

		msg := tgbotapi.NewMessage(-1001287472972, text) // 582130977

		if isFirstRun {
			poll := newPoll(-1001287472972, "Кто идет?", "Я иду", "Я остаюсь голодным")
			_, err = bot.Send(poll)
		}

		_, err = bot.Send(msg)

		if err != nil {
			fmt.Println(err)
			return "", err
		}
	}

	return "true", nil
}

func newPoll(chatID int64, question string, options ...string) tgbotapi.SendPollConfig {
	return tgbotapi.SendPollConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID: chatID,
		},
		Question:    question,
		Options:     options,
		IsAnonymous: false,
	}
}

func getRandom(in []string) string {
	rand.Seed(time.Now().Unix())
	message := in[rand.Intn(len(in))]

	return message
}
