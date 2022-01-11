package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/orpheeh/jalbv-backend/account"
)

func getAuthAdmin(c *gin.Context) {
	value, exist := c.Get("account")
	if exist {
		response, err := getAdminByAccount(value.(account.Account).ID)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "User account not found"})
		} else {
			c.IndentedJSON(http.StatusUnauthorized, response)
		}
	} else {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "User account not found"})
	}
}
