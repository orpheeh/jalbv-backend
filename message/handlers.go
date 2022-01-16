package message

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	util "github.com/orpheeh/jalbv-backend/utils"
)

type MessageData struct {
	Fullname  string
	Telephone string
	Email     string
	Message   string
}

func sendContactMessage(c *gin.Context) {
	var newFake MessageData

	if err := c.BindJSON(&newFake); err != nil {
		return
	}

	email := os.Getenv("CONTACT")
	fmt.Println(email)
	util.SendEmail(email, "Message d'internaute JALBV", contactMessageContent(newFake))

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Message envoyé !"})
}

func contactMessageContent(data MessageData) string {
	return fmt.Sprintf(`
	<p>Bonjour,</p>
	<p>Un nouveau message vient d'être envoyé depuis le site jarrivealibreville.com,</p>
	
	<p><strong>Nom et prénom:</strong> %v</p>
	<p><strong>Adresse email:</strong> %v</p>
	<p><strong>Téléphone:</strong> %v</p>
	<br>
	<strong>Message:</strong><br/>
	<p>%v</p>
	
	<br />
	<p>Pour des raisons de sécurité vous êtes invité à changer votre mot de passe régulièrement !</p>
	`, data.Fullname, data.Email, data.Telephone, data.Message)
}
