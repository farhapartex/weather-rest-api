package main

import (
	"example/weather-rest-api/config"
	"example/weather-rest-api/server"
	"flag"
	"fmt"
)

func main() {
	// Package flag implements command-line flag parsing.
	environment := flag.String("env", "development", "")
	flag.Parse()

	// system configurations
	config.Init(*environment)
	fmt.Println(config.GetConfig().GetString("server.port"))

	// server functionalities
	server.Init()
}
