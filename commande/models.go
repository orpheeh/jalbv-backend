package commande

import (
	"github.com/orpheeh/jalbv-backend/client"
	"github.com/orpheeh/jalbv-backend/produit"
)

type Commande struct {
	ID                    int             `json:"id"`
	IsImport              bool            `json:"isImport"`
	IsExport              bool            `json:"isExport"`
	PaysDepart            string          `json:"paysDepart"`
	PaysArrive            string          `json:"paysArrive"`
	ExpediteurNom         string          `json:"expediteurNom"`
	ExpediteurAdresse     string          `json:"expediteurAdresse"`
	ExpediteurBP          string          `json:"expediteurBP"`
	ExpediteurTelephone   string          `json:"expediteurTelephone"`
	ExpediteurEmail       string          `json:"expediteurEmail"`
	DestinataireNom       string          `json:"destinataireNom"`
	DestinataireAdresse   string          `json:"destinataireAdresse"`
	DestinataireEmail     string          `json:"destinataireEmail"`
	DestinataireTelephone string          `json:"destinataireTelephone"`
	DestinataireBP        string          `json:"destinataireBP"`
	PortChargement        string          `json:"portChargement"`
	PaysChargement        string          `json:"paysChargement"`
	PortDechargement      string          `json:"portDechargement"`
	PaysDechargement      string          `json:"paysDechargement"`
	PackingList           string          `json:"packingList"`
	Societe               string          `json:"societe"`
	NIF                   string          `json:"NIF"`
	ProduitId             int             `json:"produitId"`
	Etape                 int             `json:"etape"`
	Produit               produit.Produit `json:"produit"`
	ClientId              int             `json:"clientId"`
	Client                client.Client   `json:"client"`
	Date                  string          `json:"date"`
}

type Colis struct {
	ID         int    `json:"id"`
	Quantite   int    `json:"quantite"`
	Largeur    int    `json:"largeur"`
	Longueur   int    `json:"longueur"`
	Hauteur    int    `json:"hauteur"`
	Poids      int    `json:"poids"`
	PhotoURL   string `json:"photoURL"`
	CommandeId int    `json:"commandeId"`
}

type Courrier struct {
	ID         int    `json:"id"`
	Quantite   int    `json:"quantite"`
	Largeur    int    `json:"largeur"`
	Longueur   int    `json:"longueur"`
	Hauteur    int    `json:"hauteur"`
	Poids      int    `json:"poids"`
	PhotoURL   string `json:"photoURL"`
	CommandeId int    `json:"commandeId"`
	Nom        string `json:"nom"`
	Type       string `json:"type"`
}

type Conteneur struct {
	ID         int    `json:"id"`
	Type       string `json:"type"`
	CommandeId int    `json:"commandeId"`
}
