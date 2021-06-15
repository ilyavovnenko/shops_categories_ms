package main

import (
	"log"
	"time"

	config "github.com/ilyavovnenko/shops_categories_ms/configs"
	"github.com/ilyavovnenko/shops_categories_ms/init/db"
	"github.com/ilyavovnenko/shops_categories_ms/init/router"

	_userHandler "github.com/ilyavovnenko/shops_categories_ms/internal/user/delivery/http"
	_userRepo "github.com/ilyavovnenko/shops_categories_ms/internal/user/repository/mysql"
	_userUsecase "github.com/ilyavovnenko/shops_categories_ms/internal/user/usecase"

	_ "github.com/go-sql-driver/mysql"
	fiber "github.com/gofiber/fiber/v2"
	viper "github.com/spf13/viper"
)

func main() {
	conf := config.GetConfig("config.json")

	dbConnection, err := db.GetDbConnection(conf.Database)
	if err != nil {
		log.Fatal(err)
	}

	// close db connection in the end of main function processing
	defer db.CloseDbConnection(dbConnection)

	app := router.InitRouter()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the measurements API.")
	})

	userRepository := _userRepo.New(dbConnection)

	// init usecases
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	userUsecase := _userUsecase.New(userRepository, timeoutContext)
	_userHandler.New(app, userUsecase, conf.Default.User)

	app.Listen(conf.Server.Address)
}
