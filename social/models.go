package social

type Social struct {
	ID       string `json:"id"`
	ImageURL string `json:"imageURL"`
	Libelle  string `json:"libelle"`
	Link     string `json:"link"`
	IsActive string `json:"isActive"`
}
