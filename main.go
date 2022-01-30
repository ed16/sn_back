package main

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type E struct {
	Events string
}

func main() {
	r := gin.Default()

	// r.GET("/sendMessage", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	r.POST("/sendMessage", func(c *gin.Context) {
		x, _ := ioutil.ReadAll(c.Request.Body)

		c.JSON(200, gin.H{
			"message": string(x),
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
