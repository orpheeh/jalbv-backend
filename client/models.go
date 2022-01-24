package client

import "github.com/orpheeh/jalbv-backend/account"

type Client struct {
	ID         string          `json:"id"`
	Nom        string          `json:"nom"`
	Prenom     string          `json:"prenom"`
	Email      string          `json:"email"`
	Telephone  string          `json:"telephone"`
	Adresse    string          `json:"adresse"`
	Profession string          `json:"profession"`
	Entreprise string          `json:"entreprise"`
	Account    account.Account `json:"account"`
	AccountId  string          `json:"accountId"`
}
