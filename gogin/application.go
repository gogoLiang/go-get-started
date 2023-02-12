package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 需要大写
type DataResult struct {
	Code string
	Data interface{}
}

var r = gin.Default()

func main() {
	r.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, DataResult{"SUCCESS", "LST45454444"})
	})
	r.Run(":9999")
}
