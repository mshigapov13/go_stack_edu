package telegram

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	config "gitlab.ozon.dev/mshigapov13/hw/config/bots"

	// handlers "gitlab.ozon.dev/mshigapov13/hw/internal/adapters/telegram/handlers"
	botCmds "gitlab.ozon.dev/mshigapov13/hw/internal/adapters/telegram/commands"
	ports "gitlab.ozon.dev/mshigapov13/hw/internal/ports/competitors"
)

type cmdHandler func(string) string
type Bot struct {
	API         *tgbotapi.BotAPI
	name        string
	competition ports.CompetitionService
	router      map[string]cmdHandler
}

func InitTgBot(cfg config.Bot, competition ports.CompetitionService) (*Bot, error) {
	var (
		botAPI *tgbotapi.BotAPI
		err    error
	)

	botAPI, err = tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		s := fmt.Sprintf(botCreationIsFalied_format, err.Error())
		return nil, fmt.Errorf(s)
	}

	botAPI.Debug = true
	log.Printf(authorizedOnAccount_format, botAPI.Self.UserName)

	bot := &Bot{
		name:        cfg.Name,
		API:         botAPI,
		competition: competition,
		router:      make(map[string]cmdHandler),
	}
	bot.AddHandlers()
	return bot, nil
}

func (b *Bot) Run() error {
	update := tgbotapi.NewUpdate(0)
	update.Timeout = 30
	updates := b.API.GetUpdatesChan(update)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if cmd := update.Message.Command(); cmd != "" {
			if f, ok := b.router[cmd]; ok {
				msg.Text = f(update.Message.CommandArguments())
			} else {
				msg.Text = responseIfError(botCmds.ErrUnknownCommand, botCmds.AvailableCommandsTitle+"\n/"+botCmds.HelpCmd)
			}
		} else {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			msg.Text = fmt.Sprintf(userSendedTextIs_format, update.Message.Text)
		}
		_, err := b.API.Send(msg)
		if err != nil {
			log.Printf(respWasntSended_format, err)
			return err
		}
	}
	return nil
}
