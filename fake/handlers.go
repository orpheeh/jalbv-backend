package fake

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getFakes(c *gin.Context) {
	fakes, err := getAllFakes()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, fakes)
}

func createNewFake(c *gin.Context) {
	var newFake Fake

	if err := c.BindJSON(&newFake); err != nil {
		return
	}

	id, err1 := addFake(newFake)
	if err1 != nil {
		fmt.Println(err1)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"location": err1})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"location": id})
	}
}
