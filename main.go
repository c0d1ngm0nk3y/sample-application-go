package main

import (
	"net/http"

	"github.com/c0d1ngm0nk3y/sample-application-go/pkg/quotation"
	"github.com/gin-gonic/gin"
)

func handleQuote(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			ctx.String(500, "panic")
		}
	}()

	input := make([]byte, 5)
	_, err := ctx.Request.Body.Read(input)
	if nil != err {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	err = quotation.Do(input)
	if nil != err {
		ctx.String(http.StatusBadRequest, err.Error())
	}
}

func getRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.POST("/quotation", handleQuote)
	return router
}

func main() {
	r := getRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
