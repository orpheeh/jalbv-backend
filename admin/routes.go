package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/orpheeh/jalbv-backend/auth"
)

func Routing(gr *gin.Engine) {
	group := gr.Group("/admin")

	group.GET("", auth.Authorise(), getAuthAdmin)
	group.PUT("", auth.Authorise(), updateAuthAdminPassword)
}
