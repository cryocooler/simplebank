package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/cryocooler/simplebank/util"
	_ "github.com/lib/pq"
)

var TestQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("could not load config", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("could not connect to db:", err)
	}

	TestQueries = New(testDB)

	os.Exit(m.Run())

}
