package commands

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/nicogiagnoni/lolbot/riot/client"
)

func Help(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID,
		`Hi! I'm LoLSummonerBot. Here are my commands:

/m summoner_region summoner_name
/r summoner_region summoner_name
/rt summoner_region summoner_name

GL & HF :) ...`)
	msg.ReplyToMessageID = message.MessageID
	bot.Send(msg)
}

func RankedSolo(bot *tgbotapi.BotAPI, message *tgbotapi.Message, arguments []string) {
	if len(arguments) != 2 {
		Help(bot, message)
		return
	}

	ranking, er := client.GetSummonerRanking(arguments[1], arguments[0])

	var res string

	if er != nil {
		//res = "Could not execute command. Please try again"
		res = er.Error()
	} else {
		res = ranking
	}

	resMsg := tgbotapi.NewMessage(message.Chat.ID, res)
	resMsg.ReplyToMessageID = message.MessageID

	bot.Send(resMsg)
}

func SummonerMasteryChamps(bot *tgbotapi.BotAPI, message *tgbotapi.Message, arguments []string) {
	if len(arguments) != 2 {
		Help(bot, message)
		return
	}

	ranking, er := client.GetSummonerMasteryChamps(arguments[1], arguments[0])

	var res string

	if er != nil {
		res = "Could not get summoner mastery champs"
	} else {
		res = ranking
	}

	resMsg := tgbotapi.NewMessage(message.Chat.ID, res)
	resMsg.ReplyToMessageID = message.MessageID

	bot.Send(resMsg)
}

func RankedTeam(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	resMsg := tgbotapi.NewMessage(message.Chat.ID, "Todavia no funca esta parte")
	resMsg.ReplyToMessageID = message.MessageID

	bot.Send(resMsg)
}
