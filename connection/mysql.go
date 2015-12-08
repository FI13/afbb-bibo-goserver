package connection

import (
	"database/sql"
	"log"

	"github.com/fi13/afbb-bibo-goserver/model"
	"github.com/go-gorp/gorp"
)

var MySqlConnection *gorp.DbMap

func Setup() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:@/afbbbibo")
	if err != nil {
		log.Fatal("Unable to connect to mysql: %v", err)
	} else {
		MySqlConnection = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
		MySqlConnection.AddTableWithName(model.Curator{}, "benutzer")
		MySqlConnection.AddTableWithName(model.Borrower{}, "ausleiher")
	}
	return MySqlConnection
}
