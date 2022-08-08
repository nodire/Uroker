package main

import (
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	tele "gopkg.in/telebot.v3"
	"lesson-bot/service"
	"lesson-bot/system"
	"log"
	"time"
)

var (
	message, userName string
)

func main() {
	system.GetConfig("../.env")
	botConfig := tele.Settings{
		Token:  viper.GetString("BOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, checkBot := tele.NewBot(botConfig)
	if checkBot != nil {
		log.Fatalln(checkBot)
	}

	bot.Handle("/start", func(c tele.Context) error {
		return c.Send("✅ Бот запущен! Для управления бота используйте команды из список.")
	})

	bot.Handle("/previous", func(c tele.Context) error {
		userName = c.Sender().FirstName
		message = service.Scheduler.Previous(service.Request{User: userName}).Result
		return c.Send(message)
	})

	bot.Handle("/current", func(c tele.Context) error {
		userName = c.Sender().FirstName
		message = service.Scheduler.Current(service.Request{User: userName}).Result
		return c.Send(message)

	})

	bot.Handle("/next", func(c tele.Context) error {
		userName = c.Sender().FirstName
		message = service.Scheduler.Next(service.Request{User: userName}).Result
		return c.Send(message)
	})

	bot.Handle("/table", func(c tele.Context) error {
		file := &tele.Photo{
			File:    tele.FromDisk("../static/table.jpg"),
			Caption: "‼ Этот файл проста для пример. У меня пока нет табель.",
		}
		return c.Send(file)
	})

	bot.Handle("/tasks", func(c tele.Context) error {
		return c.Send("⚠ Пока этот раздел пусто.")
	})

	bot.Handle("/announcements", func(c tele.Context) error {

		return c.Send(service.Announcements())
	})

	bot.Handle("/about", func(c tele.Context) error {
		file := &tele.Photo{
			File: tele.FromDisk("../static/bot.jpg"),
			Caption: "👨‍💻 Создатель бота: Нарбаев Надир\n" +
				"📩 Электронная почта: norboyev@nodir.net\n" +
				"🌍 Сайт автора: www.nodir.net\n" +
				"📜 Лицензия: MIT License",
		}
		return c.Send(file)
	})

	bot.Start()

}
