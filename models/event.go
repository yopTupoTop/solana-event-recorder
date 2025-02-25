package models

import "database/sql"

type SolanaEvent struct {
}

func SaveEvent(db *sql.DB, event SolanaEvent) error {
	_, err := db.Exec("INSERT into ...", event)
	return err
}
