package main

import "fmt"
import "github.com/gin-gonic/gin"

func main() {
	fmt.Println("runing...")

	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	fmt.Println("exit...")
	return
}
