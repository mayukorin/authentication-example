package main

import (
	"authentication_example/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hello", handlers.GetHello)
	router.Run(":1991")
}
