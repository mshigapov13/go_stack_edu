package telegram

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	config "gitlab.ozon.dev/mshigapov13/hw/config/bots"
)

type Bot struct {
	API  *tgbotapi.BotAPI
	name string
}

func InitTgBot(cfg config.Bot) (*Bot, error) {
	var (
		botAPI *tgbotapi.BotAPI
		err    error
	)

	botAPI, err = tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		log.Fatalf(botCreationIsFalied_format, err)
	}

	botAPI.Debug = true
	log.Printf(authorizedOnAccount_format, botAPI.Self.UserName)

	bot := Bot{
		name: cfg.Name,
		API:  botAPI,
	}
	return &bot, nil
}

func (bot *Bot) Run() error {
	update := tgbotapi.NewUpdate(0)
	update.Timeout = 30
	updates := bot.API.GetUpdatesChan(update)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if cmd := update.Message.Command(); cmd != "" {
			switch cmd {
			case listCmd:
				msg.Text = listFunc()
			case createCmd:
				msg.Text = createFunc()
			case readCmd:
				msg.Text = readFunc()
			case updateCmd:
				msg.Text = updateFunc()
			case deleteCmd:
				msg.Text = deleteFunc()
			case startCmd, helpCmd:
				msg.Text = availabeCmdTitle()
			default:
				msg.Text = responseTextForWrongRequestedComand(errUnknownCommand)
			}
		} else {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			msg.Text = fmt.Sprintf(userSendedTextIs_format, update.Message.Text)
		}
		_, err := bot.API.Send(msg)
		if err != nil {
			log.Printf(msgWasntSended_format, err)
		}
	}
	return nil
}

func deleteFunc() string {
	return "was requested DELETE"
}

func updateFunc() string {
	return "was requested UPDATE"
}

func readFunc() string {
	return "was requested READ"
}

func createFunc() string {
	return "was requested CREATE"
}

func listFunc() string {
	return "was requested LIST"
}
