package client

import (
	"github.com/gin-gonic/gin"
	"github.com/orpheeh/jalbv-backend/auth"
)

func Routing(gr *gin.Engine) {

	group := gr.Group("/client")

	group.POST("", createNewClient)
	group.GET("", auth.Authorise(), getAuthClient)
	group.PUT("/:id", auth.Authorise(), editClientData)
}
