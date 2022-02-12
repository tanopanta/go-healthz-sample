package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	api := gin.Default()
	api.GET("/", o)
	api.Run(":4001")
}

func o(c *gin.Context) {
	c.String(http.StatusOK, "o")
}
