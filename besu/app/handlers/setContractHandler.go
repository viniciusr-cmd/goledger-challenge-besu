package handlers

import (
	"goledger-challenge/vinicius/besu/contract"
	"goledger-challenge/vinicius/besu/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetContractHandler(c *gin.Context) {
	var err error
	var contractValue struct {
		Value uint `json:"value"`
	}

	if err := c.ShouldBindJSON(&contractValue); err != nil {
		c.JSON(400, gin.H{"transactionStatus": "error", "error": err.Error()})
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

	var existingContract models.Contract
	if err := db.Where("address = ?", os.Getenv("CONTRACT_ADDRESS")).First(&existingContract).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			db.Save(&models.Contract{Address: os.Getenv("CONTRACT_ADDRESS"), Value: int64(contractValue.Value)})
		} else {
			c.JSON(500, gin.H{"transactionStatus": "error", "error": err.Error()})
			return
		}
	} else {
		if existingContract.Value != int64(contractValue.Value) {
			existingContract.Value = int64(contractValue.Value)
			db.Save(&existingContract)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"transactionStatus": "Success",
		"transactionMined":  receipt.Status == 1,
		"contractValue":     contractValue.Value,
		"txId":              receipt.TxHash.String(),
	})
}
