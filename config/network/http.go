package network

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/orpheeh/jalbv-backend/account"
	"github.com/orpheeh/jalbv-backend/admin"
	"github.com/orpheeh/jalbv-backend/fake"
)

func InitHttp() {
	r := gin.Default()

	fake.Routing(r)
	account.Routing(r)
	admin.Routing(r)

	r.Run(fmt.Sprintf(":%v", os.Getenv("PORT")))

}
