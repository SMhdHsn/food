package main

import (
	"log"

	"github.com/smhdhsn/food/internal/config"
	"github.com/smhdhsn/food/internal/db"
	"github.com/smhdhsn/food/internal/http"
)

// main is the main application entry.
func main() {
	conf, err := config.LoadConf()
	if err != nil {
		log.Fatal(err)
	}

	dbConn, err := db.Connect(conf.DB)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.InitMigrations(dbConn); err != nil {
		log.Fatal(err)
	}

	httpServer, err := http.New()
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Fatal(httpServer.Listen(conf.Server.Host, conf.Server.Port))
}
