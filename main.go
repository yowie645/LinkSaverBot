package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := eventconsumer.New(eventsProcessor, eventsProcessor, bathSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is not set in .env file")
	}
	return token
}
