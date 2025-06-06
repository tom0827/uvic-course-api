package main

import (
	"course-api/handlers"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	ginMode := os.Getenv("GIN_MODE")

	if ginMode == "" {
		ginMode = "debug" // Default to debug mode if not set
	}

	gin.SetMode(ginMode)

	r := gin.New()
	r.Use(gin.Logger())

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

	r.GET("/api/courses/info/:pid", handlers.InfoHandler)
	r.GET("/api/courses/sections/:term/:course", handlers.SectionHandler)
	r.GET("/api/courses", handlers.CourseHandler)

	r.Run(":8080")
}
