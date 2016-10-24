package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/nicogiagnoni/lolbot/bot/commands"
	"github.com/nicogiagnoni/lolbot/riot"
)

var bot *tgbotapi.BotAPI

func main() {
	var err error
	bot, err = tgbotapi.NewBotAPI(os.Args[1])
	riot.SetAPIKey(os.Args[2])

	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true
	log.Printf("%s Online", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil && update.InlineQuery == nil {
			continue
		}

		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			go processMessage(update.Message)
		} else if update.InlineQuery != nil {
			log.Printf("[%s] %s", update.InlineQuery.From.UserName, update.InlineQuery.Query)
			go processInlineQuery(update.InlineQuery)
		}
	}
}

func processMessage(msg *tgbotapi.Message) {
	if msg == nil {
		return
	}

	if msg.IsCommand() {
		command := msg.Command()
		arguments := strings.Split(msg.CommandArguments(), " ")

		switch command {
		case "start":
			commands.Help(bot, msg)
		case "help":
			commands.Help(bot, msg)
		case "r":
			commands.RankedSolo(bot, msg, arguments)
		case "m":
			commands.SummonerMasteryChamps(bot, msg, arguments)
		case "rt":
			commands.RankedTeam(bot, msg)
		}
	} else {
		resMsg := tgbotapi.NewMessage(msg.Chat.ID, "Sorry but I only can manage commands :)")
		resMsg.ReplyToMessageID = msg.MessageID

		bot.Send(resMsg)
	}
}

//TODO not supported yet
func processInlineQuery(query *tgbotapi.InlineQuery) {
	// response := tgbotapi.InlineConfig{
	// 	InlineQueryID: query.ID,
	// 	IsPersonal:    true,
	// 	CacheTime:     0,
	// }
	// defer func() {
	// 	_, err := bot.AnswerInlineQuery(response)
	// 	if err != nil {
	// 		log.Println("Failed to respond to query:", err)
	// 	}
	// }()
	fmt.Println(query.Query)
}
