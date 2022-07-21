package telegram

import (
	"log"

	config "gitlab.ozon.dev/mshigapov13/hw/config/bots"
	storage "gitlab.ozon.dev/mshigapov13/hw/internal/adapters/storage/inmemory"
	tgBog "gitlab.ozon.dev/mshigapov13/hw/internal/adapters/telegram/bot"
	service "gitlab.ozon.dev/mshigapov13/hw/internal/domain/figureScating/singles"
	"golang.org/x/sync/errgroup"
)

var (
	bot             *tgBog.Bot
	uneligbleCities = []string{"Moscow"}
	ageLimit        = 18
)

func Start(cfg config.Bot) {
	log.Println("Telegram bot creation attemp.")

	var err error

	db := storage.Init()
	log.Println("Storage inited")

	competition := service.Init(db, uneligbleCities, ageLimit)
	log.Println("Competition service inited")

	bot, err = tgBog.InitTgBot(cfg, competition)
	if err != nil {
		log.Fatalln("Telegram bot creation failed: ", err)
	}
	log.Println("Bot instance created")

	var g errgroup.Group
	g.Go(func() error {
		return bot.Run()
	})

	log.Println("App is started")

	err = g.Wait()
	if err != nil {
		log.Fatalln("Bot runing failed")
	}
}
