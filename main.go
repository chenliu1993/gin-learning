package main

import (
	"net/http"

	v1 "github.com/chenliu1993/gin-learning/api/v1"
	"github.com/gin-gonic/gin"
)

func main() {
	// Fake a gracefully exit
	/* stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGINT) */
	// Init a sample engine
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Simply register a router
	apiv1 := r.Group("/api/v1")
	r.Use(v1.Intercept)
	{
		apiv1.POST("/hello", v1.EchoHello)
	}
	gin.SetMode("debug")
	server := &http.Server{
		Addr:    "127.0.0.1:9090",
		Handler: r,
	}

	server.ListenAndServe()

	// <-stopCh

}
