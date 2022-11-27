package main

import (
	"flag"
	"fmt"
	mssql "golang-101-master/src/database/mssql"
	"golang-101-master/src/router/handler"
	_ "golang-101-master/src/swagger/docs"
	"net/http"
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"
)

func main() {

	port := flag.String("port", "8080", "port number")
	configPath := flag.String("config", "src/configure", "set configs path, default as: 'configure'")

	flag.Parse()
	log.Infof("port : %+v", *port)
	log.Infof("configPath directory : %+v", *configPath)
	fmt.Printf("\n[Swagger UI] API Docs at http://localhost:%s/swagger/index.html\n", *port)
	//connect database
	InitConnectionDatabase(*configPath)

	// start http server
	r := handler.Routes{}
	handleRoute := r.InitTransactionRoute()
	srv := &http.Server{
		Addr:    fmt.Sprint(":", *port),
		Handler: handleRoute,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panicf("transaction listen: %s\n", err)
		} else if err != nil {
			log.Panicf("transaction listen error: %s\n", err)
		}
		log.Infof("transaction listen at: %s", *port)
	}()

	//create channel wait signals
	//จับสัญญาณ ctr+C
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals // wait for SIGINT
}

func InitConnectionDatabase(configPath string) {
	mssql.InitDB(configPath)
}
