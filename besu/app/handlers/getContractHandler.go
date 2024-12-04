package handlers

import (
	"goledger-challenge/vinicius/besu/contract"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetContractHandler(c *gin.Context) {
	contractValue, err := contract.GetContract()
	if err != nil {
		c.JSON(400, gin.H{"transactionStatus": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transactionStatus":   "Success",
		"contractValueResult": contractValue,
	})
}
