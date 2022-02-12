package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexliesenfeld/health"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func main() {
	checker := initHealthCheck()
	healthHandler := health.NewHandler(checker)

	api := gin.Default()
	api.GET("/", helloWorld)
	api.GET("/healthz", gin.WrapH(healthHandler))

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

func initHealthCheck() health.Checker {
	return health.NewChecker(
		health.WithCacheDuration(10*time.Second),
		health.WithTimeout(30*time.Second),
		health.WithCheck(health.Check{
			Name: "w",
			Check: func(ctx context.Context) error {
				return worldCheck(ctx, "4000")
			},
		}),
		health.WithCheck(health.Check{
			Name: "o",
			Check: func(ctx context.Context) error {
				return worldCheck(ctx, "4001")
			},
		}),
		health.WithCheck(health.Check{
			Name: "r",
			Check: func(ctx context.Context) error {
				return worldCheck(ctx, "4002")
			},
		}),
		health.WithCheck(health.Check{
			Name: "l",
			Check: func(ctx context.Context) error {
				return worldCheck(ctx, "4003")
			},
		}),
		health.WithCheck(health.Check{
			Name: "d",
			Check: func(ctx context.Context) error {
				return worldCheck(ctx, "4004")
			},
		}),
		health.WithStatusListener(func(ctx context.Context, state health.CheckerState) {
			log.Println(fmt.Sprintf("health status changed to %s", state.Status))
		}),
	)
}

func worldCheck(ctx context.Context, port string) error {
	client := resty.New()

	_, err := client.R().
		SetContext(ctx). // Set Context for health check timout.
		EnableTrace().
		Get(fmt.Sprintf("http://localhost:%s", port))
	if err != nil {
		return err
	}

	return nil
}
