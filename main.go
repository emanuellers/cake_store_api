package main

import (
	"flag"

	"github.com/emanuellers/cake_store_api/storage"
)

var mode string

func init() {
	flag.StringVar(&mode, "mode", "", "Defines the mode for initialization of app.")
}

func main() {
	flag.Parse()

	switch mode {
	case "migration":
		m := storage.Migration{}
		m.Up()
		break
	}

}
