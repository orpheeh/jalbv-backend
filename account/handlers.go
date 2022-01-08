package account

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	util "github.com/orpheeh/jalbv-backend/utils"
)

func createAccount(c *gin.Context) {
	var newAccount Account
	if err := c.BindJSON(&newAccount); err != nil {
		return
	}
	newAccount.Password = util.Hash(newAccount.Password)
	id, err1 := addAccount(newAccount)
	if err1 != nil {
		fmt.Println(err1)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"location": err1})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"location": id})
	}
}

func updateAccount(c *gin.Context) {

}

func findAccountByID(c *gin.Context) {

}

func findAllAccount(c *gin.Context) {
	accounts, err := getAllCount()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, accounts)
}

func login(c *gin.Context) {
	email := c.Param("email")
	password := c.Param("password")
	account, err := getAccountByEmail(email)
	fmt.Println(err)
	if err == nil {
		isLogged := util.CompareHash(password, account.Password)
		if isLogged {
			c.IndentedJSON(http.StatusOK, gin.H{"token": "The token", "account": account})
		} else {
			c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Password error"})
		}
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Email error"})
	}
}
