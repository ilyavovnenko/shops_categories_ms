package config

import (
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Debug     bool
	Context   Context
	Database  Database
	Migration Migration
	Server    Server
	Default   Default
	Parsers   Parsers
}

type Context struct {
	Timeout int
}

type Database struct {
	Dialect  string
	Host     string
	Name     string
	Port     string
	Pass     string
	User     string
	Location string
	Try      int
}

type Migration struct {
	Folder string
}

type Server struct {
	Address string
}

type Default struct {
	PerPage uint16
}

type Categories struct {
	PerPage uint16
}

type Parsers struct {
	Amazon Amazon
	Bol    Bol
}

type Amazon struct {
	DataModelUrl string
}

type Bol struct {
	DataModelUrl string
}

func GetConfig(configPath string, log logrus.Logger) Config {
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	debugVal, err := strconv.ParseBool(viper.GetString("debug"))
	if err != nil {
		fmt.Println(err)
	}

	return Config{
		debugVal,
		getContext(log),
		getDatabase(log),
		getMigration(),
		getServer(),
		getDefault(log),
		getParsers(),
	}
}

func getDatabase(log logrus.Logger) Database {
	tryCount, err := strconv.Atoi(viper.GetString("database.try"))
	if err != nil {
		log.Error(err)
	}

	return Database{
		viper.GetString("database.dialect"),
		viper.GetString("database.host"),
		viper.GetString("database.name"),
		viper.GetString("database.port"),
		viper.GetString("database.pass"),
		viper.GetString("database.user"),
		viper.GetString("database.location"),
		tryCount,
	}
}

func getMigration() Migration {
	return Migration{
		viper.GetString("migration.folder"),
	}
}

func getContext(log logrus.Logger) Context {
	integer, err := strconv.Atoi(viper.GetString("context.timeout"))
	if err != nil {
		log.Error(err)
	}

	return Context{
		integer,
	}
}

func getServer() Server {
	return Server{
		viper.GetString("server.address"),
	}
}

func getDefault(log logrus.Logger) Default {
	var perPage uint16
	tPerPage, err := strconv.Atoi(viper.GetString("default.perPage"))
	if err != nil {
		perPage = 15

		log.Error(err)
		log.Info("Default value for default.perPage set to ", perPage)
	} else {
		perPage = uint16(tPerPage)
	}

	return Default{
		perPage,
	}
}

func getParsers() Parsers {
	return Parsers{
		getAmazon(),
		getBol(),
	}
}

func getAmazon() Amazon {
	return Amazon{
		viper.GetString("parsers.amazon.datamodel_url"),
	}
}

func getBol() Bol {
	return Bol{
		viper.GetString("parsers.bol.datamodel_url"),
	}
}
