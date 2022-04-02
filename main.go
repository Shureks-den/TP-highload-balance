package main

import (
	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

var HitsCount = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "Call_total",
	Help: "Number of calls successfully processed.",
})

func main() {
	server := echo.New()

	server.GET("/", Handle)

	err := prometheus.Register(HitsCount)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		metricsRouter := echo.New()
		metricsRouter.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
		log.Fatal(metricsRouter.Start(":8088"))
	}()

	err = server.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func Handle(c echo.Context) error{
	defer HitsCount.Add(1)
	time.Sleep(500 * time.Millisecond)
	return c.String(http.StatusOK, "I AM ALIVE")
}