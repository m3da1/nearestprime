package controller

import (
	"log"
	"net/http"
	"strconv"

	"com.m3d.dev/nearestprimes/pkg/http/rest/service"
	"github.com/gin-gonic/gin"
)

// GET /nearest/:num
// Returns the highest prime number lower than the provided number
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
