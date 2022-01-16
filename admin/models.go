package admin

import "github.com/orpheeh/jalbv-backend/account"

type Admin struct {
	ID        string `json:"id"`
	Nom       string `json:"nom"`
	Prenom    string `json:"prenom"`
	Account   account.Account
	AccountId string `json:"accountId"`
}
