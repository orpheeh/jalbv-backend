package admin

import "github.com/orpheeh/jalbv-backend/account"

type Admin struct {
	ID      int
	Nom     string
	Prenom  string
	account account.Account
}
