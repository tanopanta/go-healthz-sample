package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func main() {
	api := gin.Default()
	api.GET("/", helloWorld)
	api.GET("/healthz", healthCheck)

	api.Run(":3000")
}

func helloWorld(c *gin.Context) {
	client := resty.New()

	var message string
	for _, port := range []string{"4000", "4001", "4002", "4003", "4004"} {
		resp, err := client.R().
			EnableTrace().
			Get(fmt.Sprintf("http://localhost:%s", port))
		if err != nil {
			c.String(
				http.StatusInternalServerError,
				fmt.Sprintf("Unexpected error occurred.: %+v", err),
			)
			return
		}
		message += resp.String()
	}

	c.String(http.StatusOK, "hello "+message)
}

func healthCheck(c *gin.Context) {
	client := resty.New()

	var message string
	for _, port := range []string{"4000", "4001", "4002", "4003", "4004"} {
		resp, err := client.R().
			EnableTrace().
			Get(fmt.Sprintf("http://localhost:%s", port))
		if err != nil {
			c.String(
				http.StatusServiceUnavailable,
				fmt.Sprintf("Unexpected error occurred.: %+v", err),
			)
			return
		}
		message += resp.String()
	}

	c.String(http.StatusOK, "OK.")
}
