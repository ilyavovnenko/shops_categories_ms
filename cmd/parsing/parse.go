package main

import (
	"fmt"
	"os"

	"github.com/ilyavovnenko/shops_categories_ms/configs"

	_bp "github.com/ilyavovnenko/shops_categories_ms/internal/parser/bol"
)

func main() {
	conf := config.GetConfig("config.json")
	fmt.Println(os.Args[1])
	switch os.Args[1] {
	case "bol":
		bolParser := _bp.New(conf.Parsers.Bol.DataModelUrl)
		bolParser.ParseDatamodel()
	default:
	}
}
