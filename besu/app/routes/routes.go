package routes

import (
	"goledger-challenge/vinicius/besu/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.POST("/set", handlers.SetContractHandler)
	router.GET("/get", handlers.GetContractHandler)
}
