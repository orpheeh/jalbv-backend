package fake

import (
	"github.com/gin-gonic/gin"
)

func Routing() {
	router := gin.Default()

	router.GET("/fakes", getAllFake)
}
