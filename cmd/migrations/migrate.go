package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyavovnenko/shops_categories_ms/configs"
	"github.com/ilyavovnenko/shops_categories_ms/init/db"
	"github.com/rubenv/sql-migrate"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	conf := config.GetConfig("config.json")
	migrateCommand := migrate.Up
	if os.Args[1] == "down" {
		migrateCommand = migrate.Down
	}

	migrations := &migrate.FileMigrationSource{
		Dir: conf.Migration.Folder,
	}

	dbConnection, err := db.GetDbConnection(conf.Database)
	if err != nil {
		log.Fatal(err)
	}

	// close connection in the end of function
	defer db.CloseDbConnection(dbConnection)

	n, err := migrate.Exec(dbConnection, conf.Database.Dialect, migrations, migrateCommand)
	if err != nil {
		fmt.Printf("Error while executing migration!\n%s\n", err)
	}

	fmt.Printf("Applied %d migrations!\n", n)
}
