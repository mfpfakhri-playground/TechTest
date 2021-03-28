package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"tech-test/database"
	"tech-test/routes"
	"time"

	"github.com/spf13/viper"
)

func main() {
	readConfig()
	database.Initiation()
	serveHTTP()
}

func serveHTTP() {
	r := routes.Initiation()
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	addr := "0.0.0.0:8080"
	server := &http.Server{
		Addr:         addr,
		Handler:      r,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("could not serve http server on port 8080: %s", err.Error())
	}
}

func readConfig() {
	viper.SetConfigName("config")         // name of config file (without extension)
	viper.SetConfigType("yml")            // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {                       // Handle errors reading the config file
		panic(fmt.Errorf("Error load config: %s", err.Error()))
	}
}
