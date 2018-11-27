package server

import (
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"../../config"
	"../handlers"
)

func GinServer(conf config.VerifierConfig) {
	fmt.Println("\n")
	fmt.Println("Gin server: starting on port " + strconv.Itoa(conf.Verifier_port) + " ....")

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./frontend/dist", true)))

	r.GET("/scgetbalance", handlers.SCGetBalance)
	r.GET("/plasmabalance", handlers.GetMyPlasmaBalance)

	r.GET("/conf", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"main_account_public_key": conf.Main_account_public_key,
			"plasma_operator_address": conf.Plasma_operator_address,
			"geth_account":            conf.Geth_account,
		})
	})

	r.Run(":" + strconv.Itoa(conf.Verifier_port))
}
