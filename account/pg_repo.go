package account

import (
	"database/sql"
	"fmt"

	util "github.com/orpheeh/jalbv-backend/utils"
)

var tableName string = "Account"

func AddAccount(account Account) (int64, error) {
	datas := make(map[string]string)
	datas["email"] = account.Email
	datas["password"] = account.Password
	return util.Create(tableName, datas)
}

func GetAccountByID(idParam int) (Account, error) {
	var id, email, password string
	var lastconnexion sql.NullString
	variables := []interface{}{
		&id, &email, &password, &lastconnexion,
	}
	keys := []string{"id", "email", "password", "lastconnexion"}
	var account Account
	data, err := util.ReadOne(tableName, variables, keys, fmt.Sprintf(" WHERE id = '%d'", idParam))
	if err != nil {
		return account, err
	}
	account.ID = data["id"]
	account.Email = string(data["email"])
	account.Password = string(data["password"])
	account.LastConnexion = string(data["lastconnexion"])
	return account, err
}

func GetAccountByEmail(emailParam string) (Account, error) {
	var id, email, password string
	var lastconnexion sql.NullString
	variables := []interface{}{
		&id, &email, &password, &lastconnexion,
	}
	keys := []string{"id", "email", "password", "lastconnexion"}
	var account Account
	data, err := util.ReadOne(tableName, variables, keys, fmt.Sprintf(" WHERE email = '%v'", emailParam))
	if err != nil {
		return account, err
	}
	account.ID = data["id"]
	account.Email = string(data["email"])
	account.Password = string(data["password"])
	account.LastConnexion = string(data["lastconnexion"])
	return account, err
}

func GetAllCount() ([]Account, error) {
	var id, email, password string
	var lastconnexion sql.NullString
	variables := []interface{}{
		&id, &email, &password, &lastconnexion,
	}
	keys := []string{"id", "email", "password", "lastconnexion"}
	var accounts []Account
	datas, err := util.ReadAll(tableName, variables, keys, "")
	if err != nil {
		return accounts, err
	}
	for _, data := range datas {
		var account Account
		account.ID = data["id"]
		account.Email = string(data["email"])
		account.Password = string(data["password"])
		account.LastConnexion = string(data["lastconnexion"])
		accounts = append(accounts, account)
	}
	return accounts, err
}

func EditAccount(account Account, id int) (int64, error) {
	datas := make(map[string]string)
	datas["email"] = account.Email
	datas["password"] = account.Password
	return util.Update(tableName, datas, fmt.Sprintf(" WHERE id = %v", id))
}

func DeleteAccount(id int) (int64, error) {
	return util.Delete(tableName, fmt.Sprintf(" WHERE id = %v", id))
}
