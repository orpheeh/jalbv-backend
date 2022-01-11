package produit

import util "github.com/orpheeh/jalbv-backend/utils"

var tableName = "Produit"

func GetAllProduit() ([]Produit, error) {
	var id, libelle, description, avantages, limites, forms string
	variables := []interface{}{
		&id, &libelle, &description, &avantages, &limites, &forms,
	}
	keys := []string{"id", "libelle", "description", "avantages", "limites", "forms"}
	var produits []Produit = make([]Produit, 0)
	datas, err := util.ReadAll(tableName, variables, keys, "")
	if err != nil {
		return produits, err
	}
	for _, data := range datas {
		var produit Produit
		produit.ID = data["id"]
		produit.Libelle = string(data["libelle"])
		produit.Description = string(data["description"])
		produit.Avantages = string(data["avantages"])
		produit.Limites = string(data["limites"])
		produit.Forms = string(data["forms"])
		produits = append(produits, produit)
	}
	return produits, err
}
