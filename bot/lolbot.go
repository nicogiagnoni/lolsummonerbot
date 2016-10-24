package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var bot *LoLBot

type LoLBot struct {
	TBot *tgbotapi.BotAPI
}

func B() *LoLBot {
	if bot == nil {
		bot = &LoLBot{
			TBot: nil,
		}
	}
	return bot
}

func (l *LoLBot) SetAPI(key string) *LoLBot {
	//TODO implement err management
	bt, _ := tgbotapi.NewBotAPI(key)
	l.TBot = bt
	return l
}

func (l *LoLBot) Start() {

}

func (l *LoLBot) Stop() {

}
