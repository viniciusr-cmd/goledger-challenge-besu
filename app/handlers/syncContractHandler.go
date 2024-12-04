package handlers

import (
	"context"
	"math/big"
	"net/http"
	"time"

	"goledger-challenge/vinicius/besu/contract"
	"goledger-challenge/vinicius/besu/models"

	"github.com/gin-gonic/gin"
)

func SyncContractHandler(c *gin.Context) {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	contractValue, err := contract.GetContract()
	if err != nil {
		c.JSON(400, gin.H{"transactionStatus": "error", "error": err.Error()})
		return
	}

	contractValues, ok := contractValue.([]interface{})
	if !ok || len(contractValues) == 0 {
		c.JSON(400, gin.H{"transactionStatus": "error", "error": "invalid contract value"})
		return
	}

	var lastContractValue models.Contract
	if err := db.Last(&lastContractValue).Error; err != nil {
		c.JSON(500, gin.H{"transactionStatus": "error", "error": "failed to fetch last contract value"})
		return
	}

	lastContractValue.Value = contractValues[0].(*big.Int).Int64()
	lastContractValue.UpdatedAt = time.Now()

	if err := db.Save(&lastContractValue).Error; err != nil {
		c.JSON(500, gin.H{"transactionStatus": "error", "error": "failed to update contract value"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "Successfully synced contract value to database",
		"databaseValue": lastContractValue.Value,
		"updatedValue":  contractValues[0],
	})
}
