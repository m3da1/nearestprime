package main

import (
	"log"

	"com.m3d.dev/nearestprimes/pkg/http/rest/controller"
	"com.m3d.dev/nearestprimes/pkg/http/rest/service"
	"github.com/gin-gonic/gin"
)

// Application Entry Point
func main() {
	// Setting up api service
	service.Init()
	// Setting up gin route with default middleware
	r := gin.Default()
	// Setting mode to production
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// Registering routes
	r.GET("/nearestprime/:num", controller.NearestPrime)
	r.GET("/nearestprime", controller.HealthCheck)
	// Running Server
	log.Println("Starting application")
	r.Run()
}
