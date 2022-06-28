package server

import "example/weather-rest-api/config"

func Init() {
	config := config.GetConfig()
	port := config.GetString("server.port")
	// Initialize the router
	r := NewRouter()
	// Listen and Server in the port specified in the config file
	r.Run(config.GetString(port))
}
