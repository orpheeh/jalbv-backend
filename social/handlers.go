package social

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func findAllSocial(c *gin.Context) {
	socials, err := GetAllSocial()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, socials)
}

func editSocialData(c *gin.Context) {
	idParam, _ := c.Params.Get("id")
	var newSocial Social
	if err := c.BindJSON(&newSocial); err != nil {
		return
	}
	id, err1 := UpdateSocial(newSocial, idParam)
	if err1 != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"location": err1})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"location": id})
	}
}
