package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Success",
			"time": time.Now(),
		})
	})

	r.Run()

}
