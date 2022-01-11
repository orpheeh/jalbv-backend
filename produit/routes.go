package produit

import "github.com/gin-gonic/gin"

func Routing(gr *gin.Engine) {
	group := gr.Group("/produit")

	group.GET("", findAllProduit)
}
