package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

const (
	StartCmd     = "/start"
	WriteToBDCmd = "/save"
	DeleteBDCmd  = "/delete"
)

func (p *Processor) SendMessage(text string, chatID int64, bot *tgbotapi.BotAPI) error {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (p *Processor) CheckOnAdmin(userID, adminID int64) bool {
	if userID == adminID {
		return true
	}
	return false
}

func (p *Processor) WriteToDB(chatId int64) {
	err := p.SendMessage("Идёт запись в базу данных. Это может занять определённое время", chatId, p.Bot)
	go func() {
		err = p.Redis.Write(p.Redis.Excel.OpenFile())
		if err != nil {
			err = p.SendMessage("Ошибка файла", chatId, p.Bot)
			return
		}
		err = p.SendMessage("Запись в базу данных завершена", chatId, p.Bot)
	}()

	if err != nil {
		return
	}
}

func (p *Processor) DeleteToBD(chatId int64) {
	err := p.SendMessage("Идёт удаление", chatId, p.Bot)
	go func() {
		p.Redis.Delete()
		err = p.SendMessage("Запись в базу данных завершена", chatId, p.Bot)
	}()

	if err != nil {
		return
	}
}

func (p *Processor) FindUserByStudentID(text string, chatId int64) {
	if len(text) == 1 {
		go p.Redis.Find(c, strings.ToLower(text))
		err := p.SendMessage("Информация из базы данных: "+<-c, chatId, p.Bot)
		if err != nil {
			return
		}
		return
	}
	for _, val := range strings.Split(text, " ") {
		if val != "" {
			go p.Redis.Find(c, strings.ToLower(val))
			err := p.SendMessage("Информация из базы данных: "+<-c, chatId, p.Bot)
			if err != nil {
				return
			}
		}
	}
}
