package main

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	currentYear := time.Now().Year()
	r.GET("/age", func(c *gin.Context) {
		yearParam, err := strconv.Atoi(c.Query("year"))
		age := currentYear - yearParam
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid year",
			})
			return
		}
		c.JSON(200, gin.H{
			"age": age,
		})
	})
	r.Run()
}
