package main

import (
	"github.com/gin-gonic/gin"

	"github.com/teocci/go-gin-gorilla-ws/src"
)

func main() {
	r := gin.Default()
	hub := src.NewHub()
	go hub.Run()
	r.GET("/", func(c *gin.Context) {
		src.ServeWS(hub, c.Writer, c.Request)
	})
	_ = r.Run() // listen and serve on 0.0.0.0:8080
}
