package repo

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var openDB *sqlx.DB

func getOpenConnection() *sqlx.DB {
	if openDB == nil {
		return openConnectionToDB()
	}
	return openDB
}

func openConnectionToDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
