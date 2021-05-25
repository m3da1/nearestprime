package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET /nearest
// Get nearest prime number to a number
func NearestPrime(c *gin.Context) {
	num := c.Query("num")
	num1, err := strconv.Atoi(num)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid number provided")
		return
	}
	c.String(http.StatusOK, "%v", num1)
}
