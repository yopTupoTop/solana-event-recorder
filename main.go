package main

import (
	"log"
	"solana-events-recorder/listener"
	"solana-events-recorder/models"
)

func main() {
	db, err := models.ConnectsDB()
	if err != nil {
		log.Fatalf("Error to connect database")
	}
	defer db.Close()

	listener.ListenEvent(db)
}
