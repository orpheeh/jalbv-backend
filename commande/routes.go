package commande

import (
	"github.com/gin-gonic/gin"
	"github.com/orpheeh/jalbv-backend/auth"
)

func Routing(gr *gin.Engine) {
	group := gr.Group("/commande")

	group.GET("/by-client/:id", auth.Authorise(), findAllCommandeByClient)
	group.GET("/:id", auth.Authorise(), findCommandeByID)
	group.GET("", auth.Authorise(), findAllCommande)
	group.GET("/colis/:id", auth.Authorise(), findAllColisByCommande)
	group.GET("/courrier/:id", auth.Authorise(), findAllCourrierByCommande)
	group.GET("/conteneur/:id", auth.Authorise(), findAllConteneurByCommande)

	group.POST("", auth.Authorise(), createCommande)
	group.POST("/colis", auth.Authorise(), createColis)
	group.POST("/courrier", auth.Authorise(), createCourrier)
	group.POST("/conteneur", auth.Authorise(), createConteneur)

	group.PUT("/:id", auth.Authorise(), editCommandeData)
	group.PUT("/colis/:id", auth.Authorise(), editColisData)
	group.PUT("/courrier/:id", auth.Authorise(), editCourrierData)
	group.PUT("/conteneur/:id", auth.Authorise(), editConteneurData)

}
