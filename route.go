package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RESTFUL API
func initRouter(router *gin.Engine) {
	// for Debug
	router.GET("/v1/payments/list", GetList)
}






func GetList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"list": "",
	})
}
