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
	httpErrChan, httpTLSErrChan := rest.ServerAPI(config.APIEndpoint, config.APITLSEndpoint, dbhandler)

	select {
	case err := <- httpErrChan:
		log.Fatal("HTTP error", err)
	case err := <- httpTLSErrChan:
		log.Fatal("HTTPS error", err)
	}
}
