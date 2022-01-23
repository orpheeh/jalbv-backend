package produit

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func findAllProduit(c *gin.Context) {
	produits, err := GetAllProduit()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintln(err)})
		return
	}
	c.IndentedJSON(http.StatusOK, produits)
}
