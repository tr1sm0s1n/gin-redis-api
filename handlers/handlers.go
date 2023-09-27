package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/DEMYSTIF/gin-redis-api/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func CreateOne(c *gin.Context, r *redis.Client) {
	var newCertificate models.Certificate
	if err := c.ShouldBindJSON(&newCertificate); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	if err := r.HSet(context.Background(), newCertificate.Id, newCertificate).Err(); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newCertificate)
}

func ReadOne(c *gin.Context, r *redis.Client) {
	param := c.Param("id")
	oldCertificate := r.HGetAll(context.Background(), param).Val()
	c.IndentedJSON(http.StatusCreated, oldCertificate)
}
