package parsing

import (
	"database/sql"

	log "github.com/sirupsen/logrus"

	"github.com/ilyavovnenko/shops_categories_ms/configs"
	"github.com/ilyavovnenko/shops_categories_ms/internal/shop"

	repository "github.com/ilyavovnenko/shops_categories_ms/internal"
	bp "github.com/ilyavovnenko/shops_categories_ms/internal/parser/bol"
)

func Run(conf config.Config, dbConnection *sql.DB, log log.Logger, repoCollection repository.RepoCollection, args []string) {
	argument := args[0]

	shopId, err := repoCollection.ShopRepo.GetIdByName(argument)
	if err != nil {
		log.Error("Shop with this name is not exists!")
		return
	}

	switch argument {
	case shop.AmazonDE, shop.AmazonNL, shop.AmazonCOM:
		// todo: create logick for parsing amazon categories
	case shop.BolCom:
		bolParser := bp.New(shopId, log, conf.Parsers.Bol.DataModelUrl, repoCollection.CategoryRepo, repoCollection.AttributeRepo)
		bolParser.ParseDatamodel()
	case shop.EbayCOM, shop.EbayDE, shop.EbayNL:
		// todo: create logick for parsing ebay categories
	}
	log.Info("DONE")
}
