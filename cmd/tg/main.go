package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"tgbot/cmd/config"
	event_consumer "tgbot/internal/consumer/event-consumer"
	"tgbot/internal/event/telegram"
	"tgbot/internal/storage/excel"
	redisTg "tgbot/internal/storage/redis"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	conf := config.New()

	e := excel.New(conf.Excel.ExcelFileName)
	redis := redisTg.New(conf.Redis.RedisHost, conf.Redis.RedisPassword, conf.Redis.RedisDB, *e)
	bot, err := tgbotapi.NewBotAPI(conf.Telegram.TelegramBotAPI)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = conf.Telegram.DebugTelegramBot

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	p := telegram.New(bot, u, redis, conf.Telegram.AdminID)
	c := event_consumer.New(p)

	c.Start()

}
