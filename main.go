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
	config.Init(*environment)
	fmt.Println(config.GetConfig().GetString("server.port"))
	server.Init()

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "Hello World!")
	// })
	// r.Run()
}
