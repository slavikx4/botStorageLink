package main

import (
	"flag"
	"log"
)

const (
	HOST = "telegram.api"
)

func main() {
	//token := mustToken()

	//	tgClient := telegram.New(HOST, token)
}

func mustToken() string {
	token := flag.String("telegram-bot-token", "", "telegram bot token")
	flag.Parse()
	if *token == "" {
		log.Fatal("have not telegram bot token")
	}
	return *token
}
