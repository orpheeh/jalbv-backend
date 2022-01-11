package network

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/orpheeh/jalbv-backend/account"
	"github.com/orpheeh/jalbv-backend/admin"
	"github.com/orpheeh/jalbv-backend/fake"
	"github.com/orpheeh/jalbv-backend/produit"
	"github.com/orpheeh/jalbv-backend/social"
)

func InitHttp() {
	r := gin.Default()

	r.Static("/static", "./static")

	fake.Routing(r)
	account.Routing(r)
	admin.Routing(r)
	social.Routing(r)
	produit.Routing(r)

	r.Run(fmt.Sprintf(":%v", os.Getenv("PORT")))

}
