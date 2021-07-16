package api

import (
	"database/sql"

	config "github.com/ilyavovnenko/shops_categories_ms/configs"
	log "github.com/sirupsen/logrus"
)

func Run(conf config.Config, dbConnection *sql.DB, log log.Logger) {
	log.Info("API started")
}
