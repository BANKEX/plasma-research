package server

import (
	"../../commons/config"
	"./handlers"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GinServer(conf config.VerifierConfig) {

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	gin.SetMode(gin.ReleaseMode)

	r.GET("/etherBalance", handlers.EthereumBalance)
	r.GET("/plasmaBalance", handlers.PlasmaBalance)
	r.GET("/contractAddress", handlers.PlasmaContractAddress)

	r.Run(":" + strconv.Itoa(conf.Verifier_port))
}
