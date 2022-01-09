package account

import (
	"database/sql"
	"fmt"

	"github.com/orpheeh/jalbv-backend/config/database"
	util "github.com/orpheeh/jalbv-backend/utils"
)

var tableName string = "Account"

func addAccount(account Account) (int64, error) {
	datas := make(map[string]string)
	datas["email"] = account.Email
	datas["password"] = account.Password
	return util.Create(tableName, datas)
}

func getAccountByID(id int) (Account, error) {
	var account Account
	row := database.Postgres.QueryRow(fmt.Sprintf(`SELECT * FROM "Account" WHERE email = '%d'`, id))
	var lastconnexion sql.NullString
	if err := row.Scan(&account.ID, &account.Email, &account.Password, &account.LastConnexion); err != nil {
		if err == sql.ErrNoRows {
			return account, fmt.Errorf("accountumsById %d: no such accountum", id)
		}
		return account, fmt.Errorf("accountumsById %d: %v", id, err)
	}
	account.LastConnexion = lastconnexion.String
	return account, nil
}

func getAccountByEmail(email string) (Account, error) {
	var account Account
	row := database.Postgres.QueryRow(fmt.Sprintf(`SELECT * FROM "Account" WHERE email = '%v'`, email))
	var lastconnexion sql.NullString
	if err := row.Scan(&account.ID, &account.Email, &account.Password, &lastconnexion); err != nil {
		if err == sql.ErrNoRows {
			return account, fmt.Errorf("accountumsById %v: no such accountum", email)
		}
		return account, fmt.Errorf("accountumsById %v: %v", email, err)
	}
	account.LastConnexion = lastconnexion.String
	return account, nil
}

func getAllCount() ([]Account, error) {
	var id, email, password string
	var lastconnexion sql.NullString

	variables := []interface{}{
		&id, &email, &password, &lastconnexion,
	}

	keys := []string{"id", "email", "password", "lastconnexion"}

	var accounts []Account
	datas, err := util.ReadAll(tableName, variables, keys)

	fmt.Println(err)

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

func editAccount() {

}

func deleteAccount() {

}
