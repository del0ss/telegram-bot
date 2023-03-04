package event_consumer

import (
	"tgbot/internal/event/telegram"
)

type Consumer struct {
	Processor *telegram.Processor
}

func New(p *telegram.Processor) *Consumer {
	return &Consumer{
		Processor: p,
	}
}

func (con *Consumer) Start() {
	con.Processor.DoCmd()
}
