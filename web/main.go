package main

import (
	"github.com/gin-gonic/gin"
	"github.com/orpheeh/jalbv-backend/fake"
)

func main() {
	router := gin.Default()

	fake.Routing()

	router.Run("localhost:8080")
}
