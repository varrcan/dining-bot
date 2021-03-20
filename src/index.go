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
	}

	run := []string{
		"го!",
		"Go!",
		"Поiхали!",
		"Жратикоо",
		"На обед!",
		"Идём?",
		"Пора обедать",
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

		msg := tgbotapi.NewMessage(-1001287472972, text)

		bot.Send(msg)
	}

	return "true", nil
}

func getRandom(in []string) string {
	randomIndex := rand.Intn(len(in))
	pick := in[randomIndex]

	return pick
}
