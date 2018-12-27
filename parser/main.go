package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"parser/config"
	"parser/database"
	"parser/scraping"
)

var (
	pathToConfig = flag.String("config", "", "Path to configuration JSON file")
	numThreads   = 5
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

	wg := &sync.WaitGroup{}
	urlsChan := make(chan string, numThreads)
	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go scraping.Worker(wg, db.Conn, urlsChan)
	}

	scraping.GetProductsURLs(urlsChan)

	syscallChan := make(chan os.Signal, 1)
	signal.Notify(syscallChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-syscallChan
		log.Println("Shutting down...")
		db.Disconnect()
		os.Exit(0)
	}()
	wg.Wait()
}
