package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"status":  http.StatusOK,
		})
	})

	r.GET("/me/:id", func(c *gin.Context) {
		var id = c.Param("id")

		c.JSON(http.StatusOK, gin.H{
			"user_id": id,
		})
	})

	r.POST("/me/login", func(c *gin.Context) {
		type body struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		var b body
		c.BindJSON(&b)
		c.JSON(http.StatusOK, gin.H{
			"username": b.Username,
			"password": b.Password,
		})
	})

	r.Run(":8000")
}
