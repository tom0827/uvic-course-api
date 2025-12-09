package handlers

import (
	"course-api/src/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(c *gin.Context) {
	_, err := redis.RedisClient.Ping(c).Result()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "unhealthy", "redis": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
