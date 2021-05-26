package controller

import (
	"log"
	"net/http"
	"strconv"

	"com.m3d.dev/nearestprimes/pkg/http/rest/service"
	"github.com/gin-gonic/gin"
)

// GET /nearest/:num
//
// NearestPrime godoc
// @Summary Returns the highest prime number lower than the provided number.
// @Description Gets the highest prime number lower than the provided number.
// @Tags root
// @Accept */*
// @Param num path int true "Number"
// @Produce text/plain
// @Success 200
// @Router /nearestprime/{num} [get]
func NearestPrime(c *gin.Context) {
	// Obtain string representation of the number from the path variable
	req := c.Param("num")
	// Convert string to an integer
	num, err := strconv.Atoi(req)
	// Error handling during conversion
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid number provided")
		return
	}
	// Retrieve nearest prime
	prime := service.GetNearestPrime(num)
	log.Printf("Nearest prime for %v is %v", num, prime)
	// Response to return
	c.String(http.StatusOK, "%v", prime)
}

// HealthCheck godoc
// @Summary Show the status of service.
// @Description Gets the status of service.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /nearestprime [get]
func HealthCheck(c *gin.Context) {
	// Construct response message
	response := map[string]interface{}{
		"status": "OK",
	}
	// Response to return
	c.JSON(http.StatusOK, response)
}
