package db

import (
	"os"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	client *gorm.DB
}

func connect() *DB {
	// create DB client
	db := DB{}
	var err error
	db.client, err = gorm.Open(sqlite.Open("godo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return &db
}
