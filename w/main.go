package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	api := gin.Default()
	api.GET("/", w)
	api.Run(":4000")
}

func w(c *gin.Context) {
	c.String(http.StatusOK, "w")
}
