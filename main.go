package main

import (
	"log"
	"os"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/alpas15/tg-bot/bot"
)

func init() {
	if os.Getenv("TG_TOKEN") == "" {
		log.Print("-telegrambottoken is required")
		os.Exit(1)
	}
}

func main() {
	apiBot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot := bot.Bot{
		Bot:      apiBot,
		Handlers: bot.HandlerList(),
	}

	bot.Start()
}
