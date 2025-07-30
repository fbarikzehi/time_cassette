package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"timecassette_api/routes"
)

func main() {
	router := gin.New()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://timecassette.com", "http://localhost:3000"},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routes.RegisterRoutes(router)
	if err := router.Run(":3005"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
