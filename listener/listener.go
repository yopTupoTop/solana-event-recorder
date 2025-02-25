package listener

import (
	"database/sql"
	"log"
	"solana-events-recorder/models"

	"github.com/gorilla/websocket"
)

const (
	rpcWsUrl  = ""
	programId = ""
)

func ListenEvent(db *sql.DB) {
	connection, _, err := websocket.DefaultDialer.Dial(rpcWsUrl, nil)
	if err != nil {
		log.Fatalf("Error with websocket connection")
	}
	defer connection.Close()

	subscribenMessage := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "logSubscribe",
		"params": []interface{}{
			programId,
			map[string]interface{}{
				"mentions": []string{"EventName"},
			},
		},
		"id": 1,
	}

	err = connection.WriteJSON(subscribenMessage)
	if err != nil {
		log.Fatalf("Error to write message")
	}

	_, message, err := connection.ReadMessage()
	if err != nil {
		log.Fatalf("Error to read message")
	}

	var event models.SolanaEvent
	err = decodeEvent(message, &event)
	if err != nil {
		log.Fatalf("Error to decode event")
	}

	err = models.SaveEvent(db, event)
	if err != nil {
		log.Fatalf("Error to save event into database")
	}
}
