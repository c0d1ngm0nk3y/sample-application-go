package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/c0d1ngm0nk3y/sample-application-go/pkg/model"
	"github.com/c0d1ngm0nk3y/sample-application-go/pkg/quotation"
	"github.com/gin-gonic/gin"
)

func handleQuoteJSON(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			ctx.String(500, "panic")
		}
	}()

	inputBytes, err := ioutil.ReadAll(ctx.Request.Body)
	fmt.Printf("in: %s\n", string(inputBytes))
	if nil != err {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	input := model.StringToInputFlat(inputBytes)
	price := quotation.Do(input)
	if nil != err {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	var rawInput map[string]interface{}
	json.Unmarshal(inputBytes, &rawInput)
	rawInput["price"] = price
	output, _ := json.Marshal(rawInput)
	fmt.Printf("out: %s\n", string(output))
	ctx.JSON(http.StatusOK, rawInput)
}

func handleQuote(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			ctx.String(500, "panic")
		}
	}()

	inputBytes, err := ioutil.ReadAll(ctx.Request.Body)
	if nil != err {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	input := model.StringToInput(inputBytes)
	price := quotation.Do(input)
	if nil != err {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	result := model.Result{Price: price}
	ctx.JSON(http.StatusOK, result)
}

func getRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.POST("/quotation", handleQuote)
	router.POST("/quotation.json", handleQuoteJSON)
	return router
}

func main() {
	r := getRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
