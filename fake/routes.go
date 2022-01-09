package fake

import (
	"github.com/gin-gonic/gin"
	"github.com/orpheeh/jalbv-backend/auth"
)

func Routing(gr *gin.Engine) {
	group := gr.Group("/fakes")

	group.GET("", auth.Authorise(), getFakes)
	group.POST("", createNewFake)
}
