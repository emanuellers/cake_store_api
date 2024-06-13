package main

import (
	"flag"

	route "github.com/emanuellers/cake_store_api/app/routes"
	"github.com/emanuellers/cake_store_api/database"
	"github.com/joho/godotenv"
)

var mode string

func init() {
	flag.StringVar(&mode, "mode", "", "Defines the mode for initialization of app.")
}

func main() {
	flag.Parse()
	godotenv.Load()

	switch mode {
	case "migration":
		m := database.Migration{}
		m.Up()
		break
	case "start":
		route.Routes()
		break
	}

}
