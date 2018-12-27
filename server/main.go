package main

import (
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
	"os/signal"
	"flag"
	"os"
	"syscall"
	"time"
	"server/config"
	"server/database"
	// "server/handlers"
	"server/router"

	"github.com/fatih/color"
	"github.com/valyala/fasthttp"
)

func loggerHandlerMiddleware(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		start := time.Now()
		handler(ctx)
		log.Printf("[%s] %s, %s\n", string(ctx.Method()), ctx.URI(), time.Since(start))
	})
}
var (
	pathToConfig = flag.String("config", "", "Path to configuration JSON file")
	port = ":3000"
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

	syscallChan := make(chan os.Signal, 1)
	signal.Notify(syscallChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-syscallChan // goroutine will be frozed at here cause it will be wating until signal is received.
		log.Println("Shutting down...")
		db.Disconnect()
		os.Exit(0)
	}()

	router := router.NewRouter()
	fmt.Println(color.BlueString("STARTING SERVER AT http://localhost:3000"))
	log.Fatal(fasthttp.ListenAndServe(port, loggerHandlerMiddleware(router.Handler)))
}