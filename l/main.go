package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	api := gin.Default()
	api.GET("/", l)
	api.Run(":4003")
}

func l(c *gin.Context) {
	c.String(http.StatusOK, "l")
}
