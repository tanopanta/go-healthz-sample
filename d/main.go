package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	api := gin.Default()
	api.GET("/", hello)
	api.Run(":4004")
}

func hello(c *gin.Context) {
	c.String(http.StatusOK, "d")
}
