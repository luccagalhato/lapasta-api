package main

import (
	"flag"
	"log"

	config "lapasta/config"
	database "lapasta/database"
	utils "lapasta/internal/Utils"
	server "lapasta/server"
)

var createConfig bool
var connectionLinx *database.SQLStr
var err error

func init() {
	flag.BoolVar(&createConfig, "config", false, "create config.yaml file")
	flag.Parse()

	if createConfig {
		config.CreateConfigFile()
		return
	}

	log.Print("loading config file")
	if err := config.LoadConfig(); err != nil {
		log.Fatal(err)
	}

	log.Print("connecting sql ...")
	connectionLinx, err = database.MakeSQL(config.Yml.SQL.Host, config.Yml.SQL.Port, config.Yml.SQL.User, config.Yml.SQL.Password)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	utils.SetSQLConn(connectionLinx)	
	server.Controllers()
}
