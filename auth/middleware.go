package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	util "github.com/orpheeh/jalbv-backend/utils"
)

func Authorise() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")

		split := strings.Split(authorization, " ")

		if len(split) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"messages": "Missing access token on authorization field of request headers"})
			return
		}

		bearerToken := split[1]

		account, err := util.VerifyToken(bearerToken)

		if err == nil {
			c.Set("account", account)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"messages": fmt.Sprint(err)})
		}
	}
}
