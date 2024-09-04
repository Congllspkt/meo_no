package main

import (
	utildata "meo_no/utilData"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", getStatus)
	
	r.Run()
}

func getStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": utildata.WAITING,
	})
}
