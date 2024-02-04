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

	baseAuthGroupRouter := router.Group("/basic_auth")
	baseAuthGroupRouter.Use(middleware.BasicAuthMiddleware())
	baseAuthGroupRouter.GET("/hello", handlers.GetHello)

	tokenAuthGroupRouter := router.Group("/token_auth")
	tokenAuthGroupRouter.POST("/jwt_token", handlers.GenerateJwtToken)

	router.Run(":1991")
}
