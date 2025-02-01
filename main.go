package main

import (
	"go-project/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/scan", api.Scan)
	r.POST("/query", api.Query)
	r.Run()
}
