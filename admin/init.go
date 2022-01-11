package admin

import (
	"fmt"
	"log"
	"os"

	"github.com/orpheeh/jalbv-backend/account"
	util "github.com/orpheeh/jalbv-backend/utils"
	"github.com/zhexuany/wordGenerator"
)

type adminData struct {
	Nom    string
	Prenom string
	Email  string
}

var admins = []adminData{
	{Nom: "Nve", Prenom: "Orphée", Email: "norpheehounkponou@gmail.com"},
	{Nom: "Mve", Prenom: "Orphé", Email: "orpheenve@hotmail.com"},
}

func Init() {
	for _, admin1 := range admins {
		password := wordGenerator.GetWord(16)
		account1 := account.Account{
			Email:    admin1.Email,
			Password: util.Hash(password),
		}
		account.AddAccount(account1)
		account2, err := account.GetAccountByEmail(account1.Email)
		if err != nil {
			log.Printf("Failed to create account %v : %v", account1.Email, err)
		} else {
			admin2 := Admin{
				Nom:       admin1.Nom,
				Prenom:    admin1.Prenom,
				AccountId: account2.ID,
			}
			_, err2 := addAdmin(admin2)
			if err2 != nil {
				log.Printf("Failed to create admin %v ", admin2.Nom)
			} else {
				util.SendEmail(account1.Email, "Création de compte admin JALBV", emailHTMLContent(admin2.Prenom, account1.Email, password))
				log.Printf("Account %v are successfully created !", account1.Email)
			}
		}
	}
}

func emailHTMLContent(fullname, username, password string) string {
	return fmt.Sprintf(`
	<p>Bonjour %v,</p>
	<p>Un compte vient de vous être créé sur JALBV Admin, vous pouvez maintenant accéder à la plateforme admin en utilisant les identifiants suivants:</p>
	
	<p><strong>Adresse email:</strong> %v</p>
	<p><strong>Mot de passe:</strong> %v</p>
	
	<p>Le lien de la plateforme: <a href="%v/admin/login">%v/admin/login</a></p>
	
	<br />
	<p>Pour des raisons de sécurité vous êtes invité à changer votre mot de passe régulièrement !</p>
	`, fullname, username, password, os.Getenv("URL"), os.Getenv("URL"))
}
