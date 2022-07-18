package main

import (
	"fmt"

	config "gitlab.ozon.dev/mshigapov13/hw/config"
)

func main() {

	cfg := config.LoadConfigJson()
	fmt.Println(cfg.Bot)

}
