package main

import (
	"log"

	"github.com/ilyavovnenko/shops_categories_ms/init/db"

	_ "github.com/go-sql-driver/mysql"

	config "github.com/ilyavovnenko/shops_categories_ms/configs"
)

func main() {
	conf := config.GetConfig("config.json")

	dbConnection, err := db.GetDbConnection(conf.Database)
	if err != nil {
		log.Fatal(err)
	}

	// close db connection in the end of main function processing
	defer db.CloseDbConnection(dbConnection)

	// todo: set cobra here for calling different commands
}
