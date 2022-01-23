package produit

import (
	"fmt"
	"strconv"

	util "github.com/orpheeh/jalbv-backend/utils"
)

var tableName = "Produit"

func GetAllProduit() ([]Produit, error) {
	var id, libelle, description, avantages, limites, forms string
	var isImport, isExport bool
	variables := []interface{}{
		&id, &libelle, &description, &avantages, &limites, &forms, &isImport, &isExport,
	}
	keys := []string{"id", "libelle", "description", "avantages", "limites", "forms", "isImport", "isExport"}
	var produits []Produit = make([]Produit, 0)
	datas, err := util.ReadAll(tableName, variables, keys, "")
	if err != nil {
		return produits, err
	}
	for _, data := range datas {
		var produit Produit
		produit.ID = fmt.Sprint(data["id"])
		produit.Libelle = fmt.Sprint(data["libelle"])
		produit.Description = fmt.Sprint(data["description"])
		produit.Avantages = fmt.Sprint(data["avantages"])
		produit.Limites = fmt.Sprint(data["limites"])
		produit.Forms = fmt.Sprint(data["forms"])
		produit.IsImport, _ = strconv.ParseBool(fmt.Sprint(data["isImport"]))
		produit.IsExport, _ = strconv.ParseBool(fmt.Sprint(data["isExport"]))
		produits = append(produits, produit)
	}
	return produits, err
}
