package admin

import "github.com/orpheeh/jalbv-backend/account"

type Admin struct {
	ID        string
	Nom       string
	Prenom    string
	Account   account.Account
	AccountId string
}
