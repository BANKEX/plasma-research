package handlers

import (
	"../../../config"
	"../../../listeners/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func EthereumBalance(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Balance": storage.Balance,
	})
}

func PlasmaBalance(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Balance": "0",
	})
}

func PlasmaContractAddress(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Address": config.SmartContractAddress,
	})
}
