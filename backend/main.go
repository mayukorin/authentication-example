package main

import (
	"authentication_example/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// allows all origins
	router.Use(cors.Default())
	router.GET("/hello", handlers.GetHello)
	router.Run(":1991")
}
