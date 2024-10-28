package sql

import (
	"database/sql"
	"fmt"
	"log"
	config "pure_stitch/config"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectSql() error {
	values := config.Db_Config()
	source := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", values.Username, values.Password, values.Protocol, values.Ip_address, values.Port, values.Database)
	database, err := sql.Open("mysql", source)
	if database.Ping() != nil {
		log.Fatal(database.Ping())
	}
	if err != nil {
		log.Fatal(err)
	}
	database.SetMaxOpenConns(20)
	DB = database
	return nil
}
