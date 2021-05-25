package main

import (
	"com.m3d.dev/nearestprimes/pkg/http/rest/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	// Setting up gin route with default middleware
	r := gin.Default()

	// Setting mode to production
	// gin.SetMode(gin.ReleaseMode)

	// Registering routes
	r.GET("/nearestprime", controller.NearestPrime)

	// Running Server
	r.Run()
}
