package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handle the /fibonacci path
func fibonacciHandler(c *gin.Context) {
	// Get number parameter from query
	number, err := strconv.Atoi(c.Query("number"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	fib := calculateFibonacci(number)
	c.String(http.StatusOK, fmt.Sprintf("%d", fib))
}

func calculateFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return calculateFibonacci(n-1) + calculateFibonacci(n-2)
}
