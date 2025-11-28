package main

import (
	"course-api/handlers"
	"log"
	"os"
	"strings"

	_ "course-api/docs/swagger" // Import generated docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// CORS middleware
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

// @title Course API
// @version 1.0
// @description API for accessing course information
// @host coursesystem.app
// @BasePath /api
func main() {
	ginMode := os.Getenv("GIN_MODE")

	if ginMode == "" {
		ginMode = "debug" // Default to debug mode if not set
	}

	gin.SetMode(ginMode)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(corsMiddleware()) // Add CORS middleware

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

	r.GET("/api/info", handlers.CourseInfoHandler)
	r.GET("/api/sections", handlers.SectionHandler)
	r.GET("/api/courses", handlers.CourseHandler)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
