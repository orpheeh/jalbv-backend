package social

import "github.com/gin-gonic/gin"

func Routing(gr *gin.Engine) {
	group := gr.Group("/social")

	group.GET("", findAllSocial)
	group.PUT("/:id", editSocialData)
}
