package controller

import (
	"log"
	"net/http"
	"strconv"

	"com.m3d.dev/nearestprimes/pkg/http/rest/service"
	"github.com/gin-gonic/gin"
)

// GET /nearest
// Get nearest prime number to a number
func NearestPrime(c *gin.Context) {
	req := c.Param("num")
	num, err := strconv.Atoi(req)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid number provided")
		return
	}
	prime := service.GetNearestPrime(num)
	log.Printf("Nearest prime for %v is %v", num, prime)
	c.String(http.StatusOK, "%v", prime)
}
