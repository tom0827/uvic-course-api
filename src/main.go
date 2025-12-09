package main

import (
	"course-api/src/handlers"
	"log"
	"os"
	"strings"

	"course-api/src/redis"

	"github.com/gin-gonic/gin"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	ginMode := os.Getenv("GIN_MODE")

	if ginMode == "" {
		ginMode = "debug"
	}
	gin.SetMode(ginMode)

	redis.InitRedis()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(corsMiddleware())

	trustedProxies := os.Getenv("TRUSTED_PROXIES")
	if trustedProxies == "" {
		err := r.SetTrustedProxies(nil)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		proxies := strings.Split(trustedProxies, ",")
		err := r.SetTrustedProxies(proxies)
		if err != nil {
			log.Fatal(err)
		}
	}
	r.GET("/api/v1/health", handlers.HealthCheckHandler)

	r.GET("/api/v1/info", handlers.CourseInfoHandler)
	r.GET("/api/v1/sections", handlers.SectionHandler)
	r.GET("/api/v1/courses", handlers.CourseHandler)

	if os.Getenv("GIN_MODE") == "debug" {
		r.GET("/api/v1/redis-stats", handlers.RedisStatsHandler)
	}

	r.Run(":8080")
}
