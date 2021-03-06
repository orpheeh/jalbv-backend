package message

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/orpheeh/jalbv-backend/commande"
	util "github.com/orpheeh/jalbv-backend/utils"
)

type MessageData struct {
	Fullname  string
	Telephone string
	Email     string
	Message   string
}

type VendezVosKilos struct {
	Fullname       string
	Telephone      string
	Email          string
	Kilos          int
	DateDepart     string
	PaysDepart     string
	PaysArrive     string
	AeroportDepart string
	AeroportArrive string
}

func sendContactMessage(c *gin.Context) {
	var newFake MessageData

	if err := c.BindJSON(&newFake); err != nil {
		return
	}

	email := os.Getenv("CONTACT")
	util.SendEmail(email, "Message d'internaute JALBV", contactMessageContent(newFake))

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Message envoyé !"})
}

func sendCommandeValidationEmail(c *gin.Context) {
	socials, err := commande.GetCommandeByID(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, socials)

	email := os.Getenv("CONTACT")
	util.SendEmail(email, fmt.Sprintf(`Validation de la commande N°%v du %v`, commande.GetCommandeID(socials), socials.Date), validationCommandeEmail(socials))

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Message envoyé !"})
}

func sendVendezVosKilosMessage(c *gin.Context) {
	var newFake VendezVosKilos

	if err := c.BindJSON(&newFake); err != nil {
		return
	}

	email := os.Getenv("CONTACT")
	util.SendEmail(email, "Vendez vos kilos", vendezVosKilosMessageContent(newFake))

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

func vendezVosKilosMessageContent(data VendezVosKilos) string {
	return fmt.Sprintf(`
	<p>Bonjour,</p>
	<p>Un nouveau message vient d'être envoyé depuis le site jarrivealibreville.com,</p>
	
	<p style="text-decoration: underline"><strong>Identification de l'individu</strong></p><br>
	
	<p><strong>Nom et prénom:</strong> %v</p>
	<p><strong>Adresse email:</strong> %v</p>
	<p><strong>Téléphone:</strong> %v</p>
	<p><strong>Nombre de kilos:</strong> %v</p>
	<br>
	
	<p style="text-decoration: underline"><strong>Informations sur le voyage</strong></p><br>

	<p><strong>Date de départ:</strong> %v</p>
	<p><strong>Aeroport de départ:</strong> %v</p>
	<p><strong>Aeroport d'arrivé:</strong> %v</p>
	<p><strong>Pays de départ:</strong> %v</p>
	<p><strong>Pays de d'arrivé:</strong> %v</p>
	<br />
	
	`, data.Fullname, data.Email, data.Telephone, data.Kilos, data.DateDepart, data.AeroportDepart, data.AeroportArrive, data.PaysDepart, data.PaysArrive)
}

func validationCommandeEmail(data commande.Commande) string {

	return fmt.Sprintf(`
	<p>Bonjour,</p>
	<p>Une nouvelle commande vient d'être validé sur le site  jarrivealibreville.com, Veuillez vous connecter à l'interface d'administration pour commencer le traitement </p>
	<br />
	<p>IDENTIFICATION DU CLIENT</p>
	<p><strong>FullName: </strong> %v %v </p>
	<p><strong>Account ID: </strong> %v </p>
	<br/>
	<p>IDENTIFICATION DE LA COMMANDE</p>
	<p><strong>Number: </strong> %v </p>
	<p><strong>Type: </strong> %v </p>
	<br/>
	<p>Ceci est un message automatique, Ne pas y répondre,</p>
	<p>Cordialement,</p>
	`, data.Client.Nom, data.Client.Prenom, data.Client.Account.Email, commande.GetCommandeID(data), data.Produit.Libelle)
}
