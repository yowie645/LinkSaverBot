package main

import (
	"flag"
	"log"

	tgClient "github.com/yowie645/ReadItLaterBot/clients/telegram"
	eventconsumer "github.com/yowie645/ReadItLaterBot/consumer/event-consumer"
	"github.com/yowie645/ReadItLaterBot/events/telegram"
	"github.com/yowie645/ReadItLaterBot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	bathSize    = 100
)

func main() {

	eventsProcessor := telegram.New(tgClient.New(tgBotHost, mustToken()), files.New(storagePath))

	log.Print("service started")

	consumer := eventconsumer.New(eventsProcessor, eventsProcessor, bathSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String("token-bot", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specifed")
	}
	return *token
}
