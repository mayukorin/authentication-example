package main

import (
	"authentication_example/handlers"
	"authentication_example/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// allows all origins
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowMethods: []string{
			"GET",
			"OPTIONS",
		},
	}))

	router.Use(middleware.BasicAuthMiddleware())
	router.GET("/hello_with_basic_auth", handlers.GetHello)
	router.Run(":1991")
}
