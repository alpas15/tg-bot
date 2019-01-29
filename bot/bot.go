package bot

import (
	"log"
	"strings"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/alpas15/tg-bot/bot/handlers"
)

type Bot struct {
	Token    string
	Bot      *tgbotapi.BotAPI
	Handlers map[string]func(tgbotapi.Update) interface{}
}

func (myBot *Bot) Start() {
	log.Printf("Authorized on account %s", myBot.Bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 5
	updates, _ := myBot.Bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		go myBot.Action(update)
		log.Printf("[%s] %v %s", update.Message.From.UserName, update.Message.Chat.ID, update.Message.Text)
	}
}

func (myBot *Bot) Action(update tgbotapi.Update) {
	msg := myBot.Handlers["default"](update)
	command := strings.TrimLeft(update.Message.Text, "/")
	if strings.Contains(command, " ") {
		commandMap := strings.Split(command, " ")
		if len(commandMap) > 0 && len(commandMap[0]) > 0 {
			command = commandMap[0]
		}
	}
	if len(command) > 0 {
		if _, er := myBot.Handlers[command]; er {
			msg = myBot.Handlers[command](update)
		}
	} else {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Something went wrong. =(")
	}
	if msg != nil {
		myBot.Bot.Send(msg.(tgbotapi.Chattable))
	}
}

func HandlerList() (handlerList map[string]func(tgbotapi.Update) interface{}) {
	handlerList = make(map[string]func(tgbotapi.Update) interface{})
	handlerList["screenshot"] = handlers.Screenshot
	handlerList["hello"] = handlers.Hello
	handlerList["default"] = handlers.Default
	return
}
