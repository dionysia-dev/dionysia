package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Stream struct {
	Source string `form:"source"`
}

func onReady(c *gin.Context) {
	var stream Stream
	if err := c.ShouldBindQuery(&stream); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": err})
		return
	}

	fmt.Printf("Stream: %+v\n", stream)

	c.JSON(200, gin.H{"status": "ok"})
}
