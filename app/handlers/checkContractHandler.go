package handlers

import (
	"context"
	"log"
	"math/big"
	"net/http"
	"time"

	"goledger-challenge/vinicius/besu/contract"
	"goledger-challenge/vinicius/besu/models"

	"github.com/gin-gonic/gin"
)

func CheckContractHandler(c *gin.Context) {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	contractValue, err := contract.GetContract()
	if err != nil {
		log.Fatalf("Error getting contract value: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting contract value"})
		return
	}

	contractValues, ok := contractValue.([]interface{})
	if !ok || len(contractValues) == 0 {
		c.JSON(400, gin.H{"transactionStatus": "error", "error": "invalid contract value"})
		return
	}

	var lastContractValue models.Contract
	if err := db.Last(&lastContractValue).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving the last contract value"})
		return
	}

	isEqual := lastContractValue.Value == int64(contractValues[0].(*big.Int).Int64())

	c.JSON(http.StatusOK, gin.H{
		"contractValue": contractValues[0].(*big.Int).Int64(),
		"databaseValue": lastContractValue.Value,
		"isEqual":       isEqual,
	})
}
