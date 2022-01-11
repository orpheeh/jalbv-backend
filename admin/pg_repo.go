package admin

import (
	"fmt"
	"strconv"

	"github.com/orpheeh/jalbv-backend/account"
	util "github.com/orpheeh/jalbv-backend/utils"
)

var tableName string = "Admin"

func addAdmin(admin Admin) (int64, error) {
	datas := make(map[string]string)
	datas["nom"] = admin.Nom
	datas["prenom"] = admin.Prenom
	datas["account"] = string(admin.AccountId)
	return util.Create(tableName, datas)
}

func getAdminByAccount(accountParam string) (Admin, error) {
	var id, nom, prenom, account1 string
	variables := []interface{}{
		&id, &nom, &prenom, &account1,
	}
	keys := []string{"id", "nom", "prenom", "account"}
	var admin Admin
	data, err := util.ReadOne(tableName, variables, keys, fmt.Sprintf(" WHERE account = '%v'", accountParam))
	if err != nil {
		return admin, err
	}
	admin.ID = data["id"]
	admin.Nom = string(data["nom"])
	admin.Prenom = string(data["prenom"])
	admin.AccountId = string(data["account"])
	v, _ := strconv.Atoi(admin.AccountId)
	admin.Account, _ = account.GetAccountByID(v)
	return admin, err
}
