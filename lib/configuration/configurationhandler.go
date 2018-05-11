package configuration

import (
	"encoding/json"
	"fmt"
	"github.com/ricardo-gomez/reservations/lib/persistence/dblayer"
	"os"
)

var (
	DBTypeDefault       = dblayer.DBTYPE("mongodb")
	DBConnectionDefault = "mongodb://192.168.0.4"
	APIEndpointDefault  = "localhost:8081"
)

type ServiceConfig struct {
	Databasetype dblayer.DBTYPE `json:"databasetype"`
	DBConnection string         `json:"dbconnection"`
	APIEndpoint  string         `json:"api_endpoint"`
}

func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{
		DBTypeDefault,
		DBConnectionDefault,
		APIEndpointDefault,
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("no se encontro el archivo de configuracion, se procede con los valores default.")
		return conf, err
	}
	err = json.NewDecoder(file).Decode(&conf)
	return conf, err
}
