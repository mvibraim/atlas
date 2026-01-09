package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err := router.Run()

	if err != nil {
    panic(err)
	}
}
