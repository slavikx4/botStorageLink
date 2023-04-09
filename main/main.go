package main

import (
	tgClient "botRofl/clients/telegram"
	"botRofl/consumer/eventConsumer"
	"botRofl/events/telegram"
	"botRofl/storage/files"
	"flag"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

// 5989559234:AAEtU3FDmLQlbAKUx9-N4-73w-nn9YeO_eQ
func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Printf("servixe started")

	consumer := eventConsumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stoped", err)
	}
}

func mustToken() string {
	token := flag.String("telegram-bot-token", "", "telegram bot token")
	flag.Parse()
	if *token == "" {
		log.Fatal("have not telegram bot token")
	}
	return *token
}
