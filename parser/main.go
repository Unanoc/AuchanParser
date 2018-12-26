package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"parser/config"
	"parser/database"
)

var (
	pathToConfig = flag.String("config", "", "Path to configuration JSON file")
)

func main() {
	flag.Parse()

	configFile, err := ioutil.ReadFile(*pathToConfig)
	if err != nil {
		log.Panic(err)
	}

	configInstance := config.Config{}
	err = json.Unmarshal(configFile, &configInstance)
	if err != nil {
		log.Fatal(err)
	}

	db := database.DB{}
	if err := db.Connect(configInstance.Connection); err != nil {
		log.Panic(err)
	}

}
