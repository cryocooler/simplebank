package main

import (
	"database/sql"
	"log"

	"github.com/cryocooler/simplebank/api"
	db "github.com/cryocooler/simplebank/db/sqlc"
	"github.com/cryocooler/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("could not connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("could not start server", err)
	}
}
