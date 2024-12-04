package routes

import (
	"goledger-challenge/vinicius/besu/handlers"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func InitRoutes(router *gin.Engine, db *gorm.DB) {
	handlers.SetDB(db)

	router.POST("/set", handlers.SetContractHandler)
	router.GET("/get", handlers.GetContractHandler)
	router.GET("/sync", handlers.SyncContractHandler)
	router.GET("/check", handlers.CheckContractHandler)
}
