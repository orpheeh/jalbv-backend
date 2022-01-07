package network

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func InitHttp() {
	r := gin.New()
	r.Use(Logger())

	r.Run(":8080")
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
