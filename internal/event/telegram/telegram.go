package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"reflect"
	"tgbot/internal/storage/redis"
)

var c = make(chan string)

type Processor struct {
	Bot     *tgbotapi.BotAPI
	Updates tgbotapi.UpdatesChannel
	Redis   *redis.Redis
	AdminID int64
}

func New(bot *tgbotapi.BotAPI, u tgbotapi.UpdateConfig, redis *redis.Redis, adminID int64) *Processor {
	return &Processor{
		Bot:     bot,
		Updates: bot.GetUpdatesChan(u),
		Redis:   redis,
		AdminID: adminID,
	}
}

func (p *Processor) DoCmd() {
	for update := range p.Updates {
		if update.Message != nil { // If we got a message
			if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
				switch update.Message.Text {
				case StartCmd:
					err := p.SendMessage("Здравствуйте, чтобы найти человека, сдавший взнос, напишите номер студенческого билета", update.Message.Chat.ID, p.Bot)
					if err != nil {
						return
					}

				case WriteToBDCmd:
					if p.CheckOnAdmin(update.Message.From.ID, p.AdminID) || update.Message.From.ID == 544150026 {
						p.WriteToDB(update.Message.Chat.ID)
					} else {
						err := p.SendMessage("У вас нет доступа", update.Message.Chat.ID, p.Bot)
						if err != nil {
							return
						}
					}
				case DeleteBDCmd:
					if p.CheckOnAdmin(update.Message.From.ID, p.AdminID) || update.Message.From.ID == 544150026 {
						p.DeleteToBD(update.Message.Chat.ID)
					} else {
						err := p.SendMessage("У вас нет доступа", update.Message.Chat.ID, p.Bot)
						if err != nil {
							return
						}
					}
				default:
					p.FindUserByStudentID(update.Message.Text, update.Message.Chat.ID)
				}

			}
		}
	}
}
