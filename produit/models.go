package produit

type Produit struct {
	ID          string `json:"id"`
	Libelle     string `json:"libelle"`
	Description string `json:"description"`
	Avantages   string `json:"avantages"`
	Limites     string `json:"limites"`
	IsImport    bool   `json:"isImport"`
	IsExport    bool   `json:"isExport"`
	Forms       string `json:"forms"`
}
