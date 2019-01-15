package handlers

import (
	"net/url"
	"strings"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/alpas15/tg-bot/bot/casperjsModule"
)

func Screenshot(update tgbotapi.Update) interface{} {
	var reply interface{}
	command := strings.TrimLeft(update.Message.Text, "/")
	commandMap := strings.Split(command, " ")
	if len(commandMap) > 1 && len(commandMap[1]) > 0 {
		if _, err := url.ParseRequestURI(commandMap[1]); err == nil {
			casperjsModule.Loader(commandMap[1], "./getPage.js")
			msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, "screenshot.png")
			reply = msg
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Url is not valid =(")
			reply = msg
		}
	} else {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You are forgote type url =(")
		reply = msg
	}
	return reply
}

func Hello(update tgbotapi.Update) interface{} {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "world!")
	return msg
}

func Default(update tgbotapi.Update) interface{} {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "What are u want from me?")
	return msg
}
