package main

import (
	"fmt"
	"log"

	"com.m3d.dev/nearestprimes/pkg/http/rest/controller"
	"com.m3d.dev/nearestprimes/pkg/http/rest/service"
	"com.m3d.dev/nearestprimes/pkg/http/rest/util"
	"github.com/gin-gonic/gin"

	_ "com.m3d.dev/nearestprimes/docs/nearestprimes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Nearest Prime API
// @version 1.0
// @description This service returns the highest prime number lower the input provided by user.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
// @schemes http
// Application Entry Point
func main() {
	// Load configuration
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Unable to load configurations")
	}
	log.Printf("%+v\n", config)
	// Setting up api service
	service.InitializePrimes(config.LargestPrime)
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
	// Swagger UI
	url := fmt.Sprintf("http://%s/swagger/doc.json", config.ServerAddress)
	swaggerURL := ginSwagger.URL(url)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL))
	// Running Server
	log.Printf("Starting application @ %v\n", config.ServerAddress)
	r.Run(config.ServerAddress)
}
