package bot

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	config "gitlab.ozon.dev/mshigapov13/hw/config/bots"
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
		log.Fatalf(botCreationIsFalied_format, err)
	}

	botAPI.Debug = true
	log.Printf(authorizedOnAccount_format, botAPI.Self.UserName)

	bot := Bot{
		name:        cfg.Name,
		API:         botAPI,
		competition: competition,
		router:      make(map[string]cmdHandler),
	}
	return &bot, nil
}

func (b *Bot) RegisterRouter(cmd string, f cmdHandler) {
	b.router[cmd] = f

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
				msg.Text = respTextForWrongComand(botCmds.ErrUnknownCommand)
			}
			// switch cmd {
			// case listCmd:
			// 	msg.Text = listFunc()
			// case createCmd:
			// 	msg.Text = createFunc()
			// case readCmd:
			// 	msg.Text = readFunc()
			// case updateCmd:
			// 	msg.Text = updateFunc()
			// case deleteCmd:
			// 	msg.Text = deleteFunc()
			// case startCmd:
			// 	msg.Text = availableCmds
			// case helpCmd:
			// 	msg.Text = cmdsDescription
			// default:
			// 	msg.Text = responseTextForWrongRequestedComand(errUnknownCommand)
			// }
		} else {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			msg.Text = fmt.Sprintf(userSendedTextIs_format, update.Message.Text)
		}
		_, err := b.API.Send(msg)
		if err != nil {
			log.Printf(respWasntSended_format, err)
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
