package main

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/ilyavovnenko/shops_categories_ms/configs"
	"github.com/ilyavovnenko/shops_categories_ms/init/db"
	"github.com/ilyavovnenko/shops_categories_ms/internal/attribute"
	"github.com/ilyavovnenko/shops_categories_ms/internal/category"
	"github.com/ilyavovnenko/shops_categories_ms/internal/shop"

	bp "github.com/ilyavovnenko/shops_categories_ms/internal/parser/bol"
)

func main() {
	conf := config.GetConfig("config.json")

	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)

	dbConnection, err := db.GetDbConnection(conf.Database)
	if err != nil {
		log.Fatal(err)
	}

	// close db connection in the end of main function processing
	defer db.CloseDbConnection(dbConnection)

	// init repos
	categoryRepo := category.New(dbConnection)
	attributeRepo := attribute.New(dbConnection)
	shopRepo := shop.New(dbConnection)

	shopId, err := shopRepo.GetIdByName(os.Args[1])
	if err != nil {
		logrus.Error("Shop with this name is not exists!")
		return
	}

	switch os.Args[1] {
	case shop.AmazonDE, shop.AmazonNL, shop.AmazonCOM:
		// todo: create logick for parsing amazon categories
	case shop.BolCom:
		bolParser := bp.New(shopId, *logrus.StandardLogger(), conf.Parsers.Bol.DataModelUrl, categoryRepo, attributeRepo)
		bolParser.ParseDatamodel()
	case shop.EbayCOM, shop.EbayDE, shop.EbayNL:
		// todo: create logick for parsing ebay categories
	}
	logrus.Info("DONE")
}
