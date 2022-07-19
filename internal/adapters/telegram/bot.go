package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/mshigapov13/hw/config"
)

type TgBot struct {
	API  *tgbotapi.BotAPI
	name string
}

func NewTgBot(cfg config.Bot) (*TgBot, error) {
	var (
		botAPI *tgbotapi.BotAPI
		err    error
	)

	botAPI, err = tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		log.Fatalf("New Telegram botAPI instanse creation failed: ", err)
	}

	botAPI.Debug = true
	log.Printf("Authorized on account %s", botAPI.Self.UserName)

	bot := TgBot{
		name: cfg.Name,
		API:  botAPI,
	}
	return &bot, nil
}

func (bot *TgBot) Start() error {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.API.GetUpdatesChan(updateConfig)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		if _, err := bot.API.Send(msg); err != nil {
			log.Printf("Message wasn't sended %s", err)
		}
	}
	return nil
}
