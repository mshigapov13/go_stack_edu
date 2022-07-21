package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	tgmbot "gitlab.ozon.dev/mshigapov13/hw/config/bots"
	tgbotapp "gitlab.ozon.dev/mshigapov13/hw/internal/application/telegram"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer cancel()

	tgmbotCfg, err := tgmbot.LoadConfigJson()
	if err != nil {
		log.Fatalf("Telegram bot Config loading failed: %s", err)
	}

	go tgbotapp.Start(tgmbotCfg.Bot)
	<-ctx.Done()
}
