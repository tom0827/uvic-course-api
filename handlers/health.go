package handlers

import (
	"course-api/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheckHandler provides a simple health check for the API and Redis
func HealthCheckHandler(c *gin.Context) {
	// Check Redis connection
	_, err := redis.RedisClient.Ping(c).Result()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "unhealthy", "redis": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
