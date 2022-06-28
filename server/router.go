package server

import (
	"example/weather-rest-api/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "I'm healthy",
		})
	})

	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			weather := v1.Group("weather")
			{
				weather.GET("/", func(ctx *gin.Context) {
					query := ctx.DefaultQuery("query", "empty")
					days := ctx.DefaultQuery("days", "1")

					if query == "empty" {
						ctx.JSON(400, gin.H{
							"message": "query is required",
						})
						return
					}

					numberOfDays, _ := strconv.Atoi(days)
					status, data := controller.WeatherData(query, numberOfDays)
					ctx.Data(status, "application/json", data)
				})
			}
		}
	}

	return router
}
