package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"../../listeners/storage"

	"github.com/gin-gonic/gin"
)

var OperatorAddress string

func SCGetBalance(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Balance": storage.Balance,
	})
}

func GetMyPlasmaBalance(c *gin.Context) {
	balance := GetReqBalance(OperatorAddress, "/pbalance")
	c.JSON(http.StatusOK, gin.H{
		"Balance": balance,
	})
}

// for get req

func GetReqBalance(address, url string) string {
	type Resp struct {
		Balance string `json:balance`
	}
	var balance Resp
	resp, err := http.Get(address + url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(body, &balance)
	if err != nil {
		log.Println(err)
	}

	return balance.Balance
}
