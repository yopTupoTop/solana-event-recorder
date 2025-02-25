package listener

import (
	"encoding/json"
	"solana-events-recorder/models"
)

func decodeEvent(eventLog []byte, event *models.SolanaEvent) error {
	err := json.Unmarshal(eventLog, event)
	return err
}
