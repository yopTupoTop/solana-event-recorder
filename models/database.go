package models

import "database/sql"

const dbConnectionString = "postgres://user:password@localhost:5432/solana_events?sslmode=disable"

func ConnectsDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
