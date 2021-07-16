package main

import (
	"database/sql"
	"log"
	"strings"

	"github.com/ilyavovnenko/shops_categories_ms/cmd/api"
	"github.com/ilyavovnenko/shops_categories_ms/cmd/migrations"
	"github.com/ilyavovnenko/shops_categories_ms/cmd/parsing"
	"github.com/ilyavovnenko/shops_categories_ms/init/db"
	"github.com/ilyavovnenko/shops_categories_ms/internal/attribute"
	"github.com/ilyavovnenko/shops_categories_ms/internal/category"
	"github.com/ilyavovnenko/shops_categories_ms/internal/shop"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"

	_ "github.com/go-sql-driver/mysql"
	config "github.com/ilyavovnenko/shops_categories_ms/configs"
	repository "github.com/ilyavovnenko/shops_categories_ms/internal"
)

func main() {
	conf := config.GetConfig("config.json")

	// logger
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)

	// db connection
	dbConnection, err := db.GetDbConnection(conf.Database)
	if err != nil {
		log.Fatal(err)
	}

	// close db connection in the end of main function processing
	defer db.CloseDbConnection(dbConnection)

	// initialising Repositories
	repoCollection := repository.RepoCollection{
		CategoryRepo:  category.New(dbConnection),
		AttributeRepo: attribute.New(dbConnection),
		ShopRepo:      shop.New(dbConnection),
	}

	// cobra part
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand( // add here commands, comma separated
		getApiCmd(conf, dbConnection),
		getMigrateCmd(conf, dbConnection),
		getParseCmd(conf, dbConnection, repoCollection),
	)

	rootCmd.Execute()
}

func getMigrateCmd(conf config.Config, dbConnection *sql.DB) *cobra.Command {
	migrationValidArgs := []string{"up", "down"}

	return &cobra.Command{
		Use:       "migrate ['" + strings.Join(migrationValidArgs, "'] ['") + "']",
		Short:     "Execute migration script",
		Long:      `Read all SQL files in migrations\sql folder and exetue every file which is not registered in migrations table`,
		Args:      cobra.ExactValidArgs(1),
		ValidArgs: migrationValidArgs,
		Run: func(cmd *cobra.Command, args []string) {
			migrations.Run(conf, dbConnection, *logrus.StandardLogger(), args)
		},
	}
}

func getParseCmd(conf config.Config, dbConnection *sql.DB, repoCollection repository.RepoCollection) *cobra.Command {
	return &cobra.Command{
		Use:   "parse ['shop.domain']",
		Short: "Execute parsing categories script",
		Long:  `Parse all categories, attributes and their values for mentioned shop`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			parsing.Run(conf, dbConnection, *logrus.StandardLogger(), repoCollection, args)
		},
	}
}

func getApiCmd(conf config.Config, dbConnection *sql.DB) *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "Start api serving",
		Long:  `Api will be served until this proces will be not killed`,
		Run: func(cmd *cobra.Command, args []string) {
			api.Run(conf, dbConnection, *logrus.StandardLogger())
		},
	}
}
