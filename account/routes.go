package account

import "github.com/gin-gonic/gin"

func Routing(gr *gin.Engine) {
	group := gr.Group("/account")

	group.GET("", findAllAccount)
	group.GET("/:id", findAccountByID)
	group.GET("/login/:email/:password", login)
	group.POST("", createAccount)
	group.PUT("/:id", updateAccount)
}
