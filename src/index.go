package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/varrcan/workcalendar/europe/russia"
)

var (
	chatID int64 = -1001287472972
)

func main() {
	if _, err := Handler(); err != nil {
		log.Println("Ошибка выполнения handler:", err)
		os.Exit(1)
	}
}

// Handler запускается каждый раз Яндекс Функцией.
// Время запуска — в UTC. Переводим в Мск (UTC+3).
func Handler() (string, error) {
	now := time.Now().UTC().Add(3 * time.Hour) // MSK = UTC+3
	hour, minute := now.Hour(), now.Minute()
	weekday := now.Weekday()

	// Для отладки
	log.Printf("Handler вызван %s (MSK)", now.Format("15:04"))

	// Проверка: только по будням и не в праздники
	if weekday == time.Saturday || weekday == time.Sunday || russia.IsHoliday(now) {
		return "Сегодня выходной или праздник, ничего не делаем", nil
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		return "", fmt.Errorf("ошибка создания бота: %w", err)
	}

	switch {
	case hour == 12 && minute >= 0 && minute < 10:
		// Окно с 12:00 до 12:09 — отправляем голосование
		prompt := getRandom(prompts)
		if err := sendPoll(bot, prompt); err != nil {
			return "", err
		}
		return "Голосование отправлено", nil

	case hour == 12 && minute >= 25 && minute < 35:
		// Окно с 12:25 до 12:34 — отправляем напоминание
		runmsg := getRandom(runMsgs)
		if err := sendMessage(bot, runmsg); err != nil {
			return "", err
		}
		return "Уведомление отправлено", nil

	default:
		// В другое время — ничего не делаем
		return "Не время для уведомлений", nil
	}
}

func sendPoll(bot *tgbotapi.BotAPI, prompt string) error {
	poll := tgbotapi.SendPollConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID: chatID,
		},
		Question:    prompt,
		Options:     []string{"Я иду", "Я остаюсь голодным"},
		IsAnonymous: false,
	}
	_, err := bot.Send(poll)
	return err
}

func sendMessage(bot *tgbotapi.BotAPI, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := bot.Send(msg)
	return err
}

// Фразы-голосования
var prompts = []string{
	"Кто идет в столовку?",
	"Будет ли толпа в столовой?",
	"Готовимся к обеду!",
	"Кто проголодался?",
	"В 12:30 идем обедать!",
	"Готовность на обед?",
	"Объявляется набор на экспедицию в столовую!",
	"Время подкрепиться — кто в деле?",
	"Чей желудок громче всех голосует за еду?",
	"Чем больше нас, тем веселее обед!",
	"Собираемся на deploy в столовую — кто участвует?",
	"Не хочешь устроить рефакторинг своего желудка?",
}

// Фразы напоминаний
var runMsgs = []string{
	"Пора кушать, выдвигаемся!",
	"Го в столовую!",
	"Время обеда, пошли!",
	"Идем?",
	"Выдвигаемся на обед!",
	"Погнали на хавку!",
	"Обеденный десант, вперед!",
	"Хватит работать, пора поесть!",
	"Ну что, обедать?",
	"В столовку!",
	"Кто не пообедал — тот не поработал!",
	"Куда идем мы с Пятачком? — Правильно, НА ОБЕД!",
	"Не доводи коллегу до «голодного» состояния!",
	"Котлеты ждут не дождутся!",
	"Бросаем мышки — берём ложки!",
	"Время сделать коммит… в желудок!",
	"Пообедал — багов меньше!",
	"Желудок просит апдейт — пора на обед.",
	"Отладка голода: шаг 1 — идти в столовую.",
	"Обед — это не баг, а полезная фича!",
}

func getRandom(arr []string) string {
	rand.Seed(time.Now().UnixNano())
	return arr[rand.Intn(len(arr))]
}
