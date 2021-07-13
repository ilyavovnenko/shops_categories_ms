package migrations

import (
	"database/sql"
	"strconv"

	"github.com/ilyavovnenko/shops_categories_ms/configs"
	"github.com/rubenv/sql-migrate"

	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

func Run(conf config.Config, dbConnection *sql.DB, log log.Logger, args []string) {
	migrateCommand := migrate.Up
	if args[0] == "down" {
		migrateCommand = migrate.Down
	}

	migrations := &migrate.FileMigrationSource{
		Dir: conf.Migration.Folder,
	}

	n, err := migrate.Exec(dbConnection, conf.Database.Dialect, migrations, migrateCommand)
	if err != nil {
		log.Error(err)
	}

	log.Info("Applied " + strconv.Itoa(n) + " migrations!")
}
