package client

import (
	"fmt"
	"strconv"

	"github.com/orpheeh/jalbv-backend/account"
	util "github.com/orpheeh/jalbv-backend/utils"
)

var tableName string = "Client"

func addClient(client Client) (int64, error) {
	fmt.Println("AddClient")
	datas := make(map[string]string)
	fmt.Println(client.Account)
	fmt.Println(client.AccountId)
	datas["nom"] = client.Nom
	datas["prenom"] = client.Prenom
	datas["email"] = client.Email
	datas["telephone"] = client.Telephone
	datas["adresse"] = client.Adresse
	datas["profession"] = client.Profession
	datas["entreprise"] = client.Entreprise
	datas["account"] = string(client.AccountId)
	return util.Create(tableName, datas)
}

func getClientByAccount(accountParam string) (Client, error) {
	var id, nom, prenom, email, telephone, adresse, profession, entreprise, account1 string
	variables := []interface{}{
		&id, &nom, &prenom, &email, &telephone, &adresse, &profession, &entreprise, &account1,
	}
	keys := []string{"id", "nom", "prenom", "email", "telephone", "adresse", "profession", "entreprise", "account"}
	var client Client
	data, err := util.ReadOne(tableName, variables, keys, fmt.Sprintf(" WHERE account = '%v'", accountParam))
	if err != nil {
		return client, err
	}
	client.ID = data["id"]
	client.Nom = string(data["nom"])
	client.Prenom = string(data["prenom"])
	client.Email = string(data["email"])
	client.Telephone = string(data["telephone"])
	client.Adresse = string(data["adresse"])
	client.Profession = string(data["profession"])
	client.Entreprise = string(data["entreprise"])

	client.AccountId = string(data["account"])
	v, _ := strconv.Atoi(client.AccountId)
	client.Account, _ = account.GetAccountByID(v)
	return client, err
}

func UpdateClient(client Client, id string) (int64, error) {
	datas := make(map[string]string)
	if client.Nom != "" {
		datas["Nom"] = client.Nom
	}
	if client.Prenom != "" {
		datas["prenom"] = fmt.Sprintf("%v", client.Prenom)
	}
	if client.Email != "" {
		datas["email"] = client.Email
	}
	if client.Telephone != "" {
		datas["telephone"] = client.Telephone
	}
	if client.Adresse != "" {
		datas["adresse"] = client.Adresse
	}
	if client.Profession != "" {
		datas["profession"] = client.Profession
	}
	if client.Entreprise != "" {
		datas["entreprise"] = client.Entreprise
	}
	return util.Update(tableName, datas, fmt.Sprintf(" WHERE id = %v", id))
}
