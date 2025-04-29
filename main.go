package main

import (
	"flag"
	"log"

	"github.com/yowie645/ReadItLaterBot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(tgBotHost, mustToken())
	//fetcher = fetcher.New(tgClient)
	//processor = processor.New(tgClient)
	//consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String("token-bot", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specifed")
	}
	return *token
}
