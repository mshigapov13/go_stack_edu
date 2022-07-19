package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"gitlab.ozon.dev/mshigapov13/hw/config"
	tgbotapp "gitlab.ozon.dev/mshigapov13/hw/internal/application/telegram"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer cancel()

	cfg, err := config.LoadConfigJson()
	if err != nil {
		log.Fatalf("Config loading failed: ", err)
	}

	go tgbotapp.Start(cfg.Bot)
	<-ctx.Done()
}
