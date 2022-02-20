package commande

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

/** UPLOAD */

func uploadPhotoColis(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	out, err := os.Create("static/upload/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	filepath := os.Getenv("API_URL") + "/static/upload/" + filename

	colis, err := GetColisByID(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	colis.PhotoURL = filepath
	fmt.Println(colis.PhotoURL)
	_, err = UpdateColis(colis, c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"filepath": filepath})
}

/** CREATE */

func createCommande(c *gin.Context) {
	var newCommande Commande
	if err := c.BindJSON(&newCommande); err != nil {
		fmt.Println(err)
		return
	}
	id, err1 := AddCommande(newCommande)
	fmt.Println(err1)
	if err1 != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"location": err1})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"location": id})
	}
}

func createColis(c *gin.Context) {
	var newCommande Colis
	if err := c.BindJSON(&newCommande); err != nil {
		return
	}
	id, err1 := AddColis(newCommande)
	if err1 != nil {
		fmt.Println(err1)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"location": err1})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"location": id})
	}
}

func createCourrier(c *gin.Context) {
	var newCommande Courrier
	if err := c.BindJSON(&newCommande); err != nil {
		return
	}
	id, err1 := AddCourrier(newCommande)
	if err1 != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"location": err1})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"location": id})
	}
}

func createConteneur(c *gin.Context) {
	var newCommande Conteneur
	if err := c.BindJSON(&newCommande); err != nil {
		return
	}
	id, err1 := AddConteneur(newCommande)
	if err1 != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"location": err1})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"location": id})
	}
}

/** DELETE */

func deleteColis(c *gin.Context) {
	idParam, _ := c.Params.Get("id")
	_, err1 := DeleteColis(idParam)
	if err1 != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprint(err1)})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "deleted successfully !"})
	}
}

func deleteCourrier(c *gin.Context) {
	idParam, _ := c.Params.Get("id")
	_, err1 := DeleteCourrier(idParam)
	if err1 != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprint(err1)})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "deleted successfully !"})
	}
}

func deleteConteneur(c *gin.Context) {
	idParam, _ := c.Params.Get("id")
	_, err1 := DeleteConteneur(idParam)
	if err1 != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprint(err1)})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "deleted successfully !"})
	}
}

/** UPDATE */

func editCommandeData(c *gin.Context) {
	idParam, _ := c.Params.Get("id")
	var newCommande Commande
	if err := c.BindJSON(&newCommande); err != nil {
		return
	}
	id, err1 := UpdateCommande(newCommande, idParam)
	if err1 != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"location": err1})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"location": id})
	}
}

func editColisData(c *gin.Context) {
	idParam, _ := c.Params.Get("id")
	var newCommande Colis
	if err := c.BindJSON(&newCommande); err != nil {
		return
	}
	id, err1 := UpdateColis(newCommande, idParam)
	if err1 != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"location": err1})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"location": id})
	}
}

func editCourrierData(c *gin.Context) {
	idParam, _ := c.Params.Get("id")
	var newCommande Courrier
	if err := c.BindJSON(&newCommande); err != nil {
		return
	}
	id, err1 := UpdateCourrier(newCommande, idParam)
	if err1 != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"location": err1})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"location": id})
	}
}

func editConteneurData(c *gin.Context) {
	idParam, _ := c.Params.Get("id")
	var newCommande Conteneur
	if err := c.BindJSON(&newCommande); err != nil {
		return
	}
	id, err1 := UpdateConteneur(newCommande, idParam)
	if err1 != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"location": err1})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"location": id})
	}
}

/** GET ALL */

func findAllCommandeByClient(c *gin.Context) {
	socials, err := GetAllCommandeByClient(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, socials)
}

func findAllCommande(c *gin.Context) {
	socials, err := GetAllCommande()
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, socials)
}

func findAllColisByCommande(c *gin.Context) {
	id := c.Param("id")
	socials, err := GetAllByCommande(id)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, socials)
}

func findAllCourrierByCommande(c *gin.Context) {
	socials, err := GetAllCourrierByCommande(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, socials)
}

func findAllConteneurByCommande(c *gin.Context) {
	socials, err := GetAllConteneurByCommande(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, socials)
}

/** GET ONE */

func findCommandeByID(c *gin.Context) {
	socials, err := GetCommandeByID(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, socials)
}
