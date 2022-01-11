package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func getAuthAdmin(c *gin.Context) {
	value, exist := c.Get("account")
	if exist {
		claims := value.(jwt.MapClaims)
		fmt.Println(claims)
		id := claims["ID"].(string)
		response, err := getAdminByAccount(id)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "User account not found"})
		} else {
			c.IndentedJSON(http.StatusOK, response)
		}
	} else {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "User account not found"})
	}
}
