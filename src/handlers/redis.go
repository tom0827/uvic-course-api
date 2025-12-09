package handlers

import (
	"course-api/src/redis"

	"github.com/gin-gonic/gin"
)

func RedisStatsHandler(c *gin.Context) {
	info, err := redis.RedisClient.Info(c, "all").Result()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.String(200, info)
}
