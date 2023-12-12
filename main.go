package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tr1sm0s1n/gin-redis-api/config"
	"github.com/tr1sm0s1n/gin-redis-api/handlers"
	"github.com/tr1sm0s1n/gin-redis-api/middlewares"
)

func main() {
	client := config.ConnectRedis()

	router := gin.Default()
	router.Use(middlewares.Authority())
	router.POST("/create", func(ctx *gin.Context) {
		handlers.CreateOne(ctx, client)
	})
	router.GET("/read/:id", func(ctx *gin.Context) {
		handlers.ReadOne(ctx, client)
	})
	router.PUT("/update/:id", func(ctx *gin.Context) {
		handlers.UpdateOne(ctx, client)
	})
	router.DELETE("/delete/:id", func(ctx *gin.Context) {
		handlers.DeleteOne(ctx, client)
	})
	router.Run("localhost:8080")
}
