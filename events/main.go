package main

import (
	"flag"
	"github.com/ricardo-gomez/reservations/events/rest"
	"github.com/ricardo-gomez/reservations/lib/configuration"
	"github.com/ricardo-gomez/reservations/lib/persistence/dblayer"
	"log"
)

func main() {
	confpath := flag.String("conf", `.\configuration\config.json`, "flag para establecer el archivo de configuracion")
	flag.Parse()
	config, _ := configuration.ExtractConfiguration(*confpath)
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	log.Fatal(rest.ServerAPI(config.APIEndpoint, dbhandler))
}
