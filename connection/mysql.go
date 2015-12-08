package connection

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	"github.com/fi13/afbb-bibo-goserver/model"
	"log"
)

var MySqlConnection *gorp.DbMap

func Setup() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:@/afbbbibo")
	if err != nil {
		log.Fatal("Unable to connect to mysql: %v", err)
	} else {
		MySqlConnection = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
		MySqlConnection.AddTableWithName(model.Curator{}, "benutzer")
	}
	return MySqlConnection
}
