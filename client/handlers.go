package client

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func createNewClient(c *gin.Context) {
	var newFake Client

	if err := c.BindJSON(&newFake); err != nil {
		fmt.Println(err)
		return
	}

	id, err1 := addClient(newFake)
	if err1 != nil {
		fmt.Println(err1)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"location": err1})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"location": id})
	}
}

func editClientData(c *gin.Context) {
	idParam, _ := c.Params.Get("id")
	var newClient Client
	if err := c.BindJSON(&newClient); err != nil {
		return
	}
	id, err1 := UpdateClient(newClient, idParam)
	if err1 != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"location": err1})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"location": id})
	}
}

func getAuthClient(c *gin.Context) {
	value, exist := c.Get("account")
	if exist {
		claims := value.(jwt.MapClaims)
		fmt.Println(claims)
		id := claims["id"].(string)
		response, err := getClientByAccount(id)
		fmt.Println(err)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "User account not found"})
		} else {
			c.IndentedJSON(http.StatusOK, response)
		}
	} else {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "User account not found"})
	}
}
