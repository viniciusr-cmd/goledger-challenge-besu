package handlers

import (
	"goledger-challenge/vinicius/besu/contract"

	"github.com/gin-gonic/gin"
)

func GetContractHandler(c *gin.Context) {
	contractValue, err := contract.GetContract()
	if err != nil {
		c.JSON(400, gin.H{"transactionStatus": "error", "error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"transactionStatus":   "Success",
		"contractValueResult": contractValue,
	})
}
