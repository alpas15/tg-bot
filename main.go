package main

import (
	"log"
	"os"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/alpas15/tg-bot/bot"
	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
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
