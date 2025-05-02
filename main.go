package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	tgClient "github.com/yowie645/ReadItLaterBot/clients/telegram"
	eventconsumer "github.com/yowie645/ReadItLaterBot/consumer/event-consumer"
	"github.com/yowie645/ReadItLaterBot/events/telegram"
	"github.com/yowie645/ReadItLaterBot/storage/sqlite"
)

const (
	tgBotHost         = "api.telegram.org"
	sqliteStoragePath = "data/sqlite/storage.db"
	bathSize          = 100
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	/* s := files.New(storagePath) */
	s, err := sqlite.New(sqliteStoragePath)
	if err != nil {
		log.Fatalf("can't connect to storage: %v", err)
	}

	if err := s.Init(context.TODO()); err != nil {
		log.Fatal("can't init storage:", err)
	}

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		s,
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
