package api

import (
	"database/sql"
	"time"

	"github.com/ilyavovnenko/shops_categories_ms/init/router"

	attributeHandler "github.com/ilyavovnenko/shops_categories_ms/internal/attribute/delivery/http"
	attributeRepository "github.com/ilyavovnenko/shops_categories_ms/internal/attribute/repository/mysql"
	attributeUsecase "github.com/ilyavovnenko/shops_categories_ms/internal/attribute/usecase"

	categoryHandler "github.com/ilyavovnenko/shops_categories_ms/internal/category/delivery/http"
	categoryRepository "github.com/ilyavovnenko/shops_categories_ms/internal/category/repository/mysql"
	categoryUsecase "github.com/ilyavovnenko/shops_categories_ms/internal/category/usecase"

	fiber "github.com/gofiber/fiber/v2"
	config "github.com/ilyavovnenko/shops_categories_ms/configs"
	repository "github.com/ilyavovnenko/shops_categories_ms/internal"
	log "github.com/sirupsen/logrus"
)

func Run(conf config.Config, dbConnection *sql.DB, log log.Logger, repoCollection repository.RepoCollection) {
	app := router.InitRouter()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the shops caegories API.")
	})

	timeoutContext := time.Duration(conf.Context.Timeout) * time.Second

	// init usecases
	categoryUsecase := categoryUsecase.New(categoryRepository.New(dbConnection, &log))
	categoryHandler.New(app, categoryUsecase, conf.Default, timeoutContext)

	attributeUsecase := attributeUsecase.New(attributeRepository.New(dbConnection, &log))
	attributeHandler.New(app, attributeUsecase, conf.Default, timeoutContext)

	app.Listen(conf.Server.Address)
}
