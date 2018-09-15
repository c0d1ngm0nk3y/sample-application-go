package main

import "github.com/gin-gonic/gin"

func getRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func main() {
	r := getRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
