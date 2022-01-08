package account

import (
	"database/sql"
	"fmt"

	"github.com/orpheeh/jalbv-backend/config/database"
)

func addAccount(account Account) (int64, error) {
	result, err := database.Postgres.Exec(fmt.Sprintf(`INSERT INTO "Account" (email, password) VALUES ('%v', '%v')`, account.Email, account.Password))
	if err != nil {
		fmt.Println(err)
		return 0, fmt.Errorf("addAccount: %v", err)
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("addAccount: %v", err)
	}
	return id, nil
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
	var accounts []Account
	rows, err := database.Postgres.Query(`SELECT * FROM "Account"`)
	if err != nil {
		return nil, fmt.Errorf("accounts : %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var account Account
		if err := rows.Scan(&account.ID, &account.Email, &account.Password); err != nil {
			return nil, fmt.Errorf("accounts: %v", err)
		}
		accounts = append(accounts, account)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("accounts: %v", err)
	}
	return accounts, nil
}

func editAccount() {

}

func deleteAccount() {

}
