package main

import (
	"flag"
	"strconv"

	"./handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	port := flag.Int("p", 8080, "port")
	flag.Parse()

	r := gin.Default()

	r.POST("/settx/:tx", handlers.SetTx)
	r.GET("/gettx/:tx", handlers.GetTx)
	r.GET("/getall", handlers.GetAllTx)
	r.GET("/publishblock", handlers.PublishBlock)
	r.GET("/pbalance", handlers.PBalance)
	r.Run(":" + strconv.Itoa(*port))
}
