package main

import (
	"log"
	"strings"

	"github.com/ilyavovnenko/shops_categories_ms/cmd/migrations"
	"github.com/ilyavovnenko/shops_categories_ms/cmd/parsing"
	"github.com/ilyavovnenko/shops_categories_ms/init/db"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"

	_ "github.com/go-sql-driver/mysql"
	config "github.com/ilyavovnenko/shops_categories_ms/configs"
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

	// cobra part
	migrationValidArgs := []string{"up", "down"}
	var cmdMigrate = &cobra.Command{
		Use:       "migrate ['" + strings.Join(migrationValidArgs, "'] ['") + "']",
		Short:     "Execute migration script",
		Long:      `Read all SQL files in migrations\sql folder and exetue every file which is not registered in migrations table`,
		Args:      cobra.ExactValidArgs(1),
		ValidArgs: migrationValidArgs,
		Run: func(cmd *cobra.Command, args []string) {
			migrations.Run(conf, dbConnection, *logrus.StandardLogger(), args)
		},
	}

	var cmdParse = &cobra.Command{
		Use:   "parse ['shop.domain']",
		Short: "Execute parsing categories script",
		Long:  `Parse all categories, attributes and their values for mentioned shop`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			parsing.Run(conf, dbConnection, *logrus.StandardLogger(), args)
		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand( // add here commands, comma separated
		cmdMigrate,
		cmdParse,
	)

	rootCmd.Execute()
}
