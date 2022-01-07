package network

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/orpheeh/jalbv-backend/fake"
)

func InitHttp() {
	r := gin.New()

	r.Use(Logger())

	fake.Routing(r)

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
