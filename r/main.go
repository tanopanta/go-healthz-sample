package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	api := gin.Default()
	api.GET("/", r)
	api.Run(":4002")
}

func r(c *gin.Context) {
	c.String(http.StatusOK, "r")
}
