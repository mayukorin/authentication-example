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
	router.Use(cors.Default())

	router.Use(middleware.BasicAuthMiddleware())
	router.GET("/hello_with_basic_auth", handlers.GetHello)
	router.Run(":1991")
}
