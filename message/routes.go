package message

import (
	"github.com/gin-gonic/gin"
)

func Routing(gr *gin.Engine) {

	group := gr.Group("/message")

	group.POST("/contact", sendContactMessage)
}
