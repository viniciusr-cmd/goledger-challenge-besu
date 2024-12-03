// FILE: handlers/set_contract_handler.go
package handlers

import (
	"goledger-challenge/vinicius/besu/contract"

	"github.com/gin-gonic/gin"
)

func SetContractHandler(c *gin.Context) {
	var err error
	var contractValue struct {
		Value uint `json:"value"`
	}

	if err := c.ShouldBindJSON(&contractValue); err != nil {
		c.JSON(400, gin.H{"transactionStatus": "error", "error": "Invalid input, please provide a valid positive integer value"})
		return
	}

	if contractValue.Value <= 0 {
		c.JSON(400, gin.H{"transactionStatus": "error", "error": "Value must be a positive integer"})
		return
	}

	receipt, err := contract.SetContract(contractValue.Value)
	if err != nil {
		c.JSON(500, gin.H{"transactionStatus": "error", "error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"transactionStatus": "Success",
		"transactionMined":  receipt.Status == 1,
		"contractValue":     contractValue.Value,
		"tx":                receipt.TxHash.String(),
	})
}
