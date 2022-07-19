package telegram

import (
	"log"

	"gitlab.ozon.dev/mshigapov13/hw/config"
	telegram "gitlab.ozon.dev/mshigapov13/hw/internal/adapters/telegram"
	"golang.org/x/sync/errgroup"
)

var bot *telegram.TgBot

func Start(cfg config.Bot) {
	log.Println("Telegram bot creation attemp.")

	var err error

	bot, err = telegram.NewTgBot(cfg)
	if err != nil {
		log.Fatalln("Telegram bot creation failed: ", err)
	}
	log.Println("Bot instance created")

	var g errgroup.Group
	g.Go(func() error {
		return bot.Start()
	})

	log.Println("App is started")

	err = g.Wait()
	if err != nil {
		log.Fatalln("Bot start failed")
	}
}
