package main

import (
	"bot/internal/config"
	"bot/internal/db"
	"bot/internal/telegram"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	cfg := config.Config{}
	if err := config.Load(&cfg); err != nil {
		fmt.Println("Error load config")
		fmt.Println(err)
	}

	var cdnDB string
	if cfg.Debag == true {
		cdnDB = cfg.Postgres.URlDBDEBAG
	} else {
		cdnDB = cfg.Postgres.UrlDB
	}

	dbstr, err := db.ConnectDB(cdnDB)
	if err != nil {
		log.Println(err)
	}
	defer dbstr.Close()

	fmt.Println(dbstr)

	telegram.StartTelegramBot(
		cfg.Tg.Api_key,
		dbstr,
	)

}
