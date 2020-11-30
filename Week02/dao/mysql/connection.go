// solmyr

package mysql

import (
	"database/sql"
	"fmt"
	"log"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	User 		string
	Password 	string
	Server 		string
	Port 		string
	Database 	string
}

func get() *sql.DB {
	var config DBConfig
	config.User = "ptsolmyr"
	config.Password = "ptsolmyr"
	config.Server = "localhost"
	config.Port = "3306"
	config.Database = "book-service"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=false&autocommit=true", config.User, config.Password, config.Server, config.Port, config.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return db
}
