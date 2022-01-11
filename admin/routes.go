package admin

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/orpheeh/jalbv-backend/auth"
)

func Routing(gr *gin.Engine) {
	recreate := os.Getenv("INIT_ADMIN")
	fmt.Println(recreate)
	if recreate == "true" {
		Init()
	}

	group := gr.Group("/admin")

	group.GET("", auth.Authorise(), getAuthAdmin)
}
