package fake

import "github.com/gin-gonic/gin"

func Routing(gr *gin.Engine) {
	group := gr.Group("/fakes")

	group.GET("", getFakes)
	group.POST("", createNewFake)
}
