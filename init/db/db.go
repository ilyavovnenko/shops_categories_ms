package db

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	config "github.com/ilyavovnenko/shops_categories_ms/configs"
)

// connect to DB several amount of times
func connectToDb(tryCount int, dialect string, credentials string) (*sql.DB, error) {
	connection, err := sql.Open(dialect, credentials)
	for i := 0; i < tryCount && err != nil; i++ {
		connection, err = sql.Open(dialect, credentials)
	}
	if err != nil {
		return connection, err
	}

	err = connection.Ping()
	for i := 0; i < tryCount && err != nil; i++ {
		err = connection.Ping()
	}

	return connection, err
}

func GetDbConnection(db config.Database) (*sql.DB, error) {
	mainCredentials := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db.User, db.Pass, db.Host, db.Port, db.Name)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", db.Location)

	credentials := fmt.Sprintf("%s?%s", mainCredentials, val.Encode())
	connection, err := connectToDb(db.Try, db.Dialect, credentials)
	if err != nil {
		return nil, err
	}

	return connection, err
}

func CloseDbConnection(dbConn *sql.DB) {
	err := dbConn.Close()
	if err != nil {
		log.Fatal(err)
	}
}
