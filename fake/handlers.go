package fake

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var fakes = []Fake{
	{ID: 1, Name: "Yoyo"},
	{ID: 2, Name: "David"},
}

func getFakes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, fakes)
}
