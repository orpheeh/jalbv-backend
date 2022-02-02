package commande

import (
	"fmt"
	"strconv"

	"github.com/orpheeh/jalbv-backend/produit"
	util "github.com/orpheeh/jalbv-backend/utils"
)

/** Commande */

func getCommandeData(commande Commande) map[string]string {
	datas := make(map[string]string)
	datas["isImport"] = fmt.Sprint(commande.IsImport)
	datas["isExport"] = fmt.Sprint(commande.IsExport)
	datas["produitId"] = fmt.Sprint(commande.ProduitId)

	datas["expediteurNom"] = fmt.Sprint(commande.ExpediteurNom)
	datas["expediteurAdresse"] = fmt.Sprint(commande.ExpediteurAdresse)
	datas["expediteurBP"] = fmt.Sprint(commande.ExpediteurBP)
	datas["expediteurEmail"] = fmt.Sprint(commande.ExpediteurEmail)
	datas["expediteurTelephone"] = fmt.Sprint(commande.ExpediteurTelephone)

	datas["destinataireNom"] = fmt.Sprint(commande.DestinataireNom)
	datas["destinataireAdresse"] = fmt.Sprint(commande.DestinataireAdresse)
	datas["destinataireBP"] = fmt.Sprint(commande.DestinataireBP)
	datas["destinataireEmail"] = fmt.Sprint(commande.DestinataireEmail)
	datas["destinataireTelephone"] = fmt.Sprint(commande.DestinataireTelephone)

	datas["portChargement"] = fmt.Sprint(commande.PortChargement)
	datas["paysChargement"] = fmt.Sprint(commande.PaysChargement)
	datas["portDechargement"] = fmt.Sprint(commande.PortDechargement)
	datas["paysDechargement"] = fmt.Sprint(commande.PaysDechargement)
	datas["packingList"] = fmt.Sprint(commande.PackingList)

	datas["societe"] = fmt.Sprint(commande.Societe)
	datas["NIF"] = fmt.Sprint(commande.NIF)
	datas["etape"] = fmt.Sprint(commande.Etape)
	datas["paysDepart"] = fmt.Sprint(commande.PaysDepart)
	datas["paysArrive"] = fmt.Sprint(commande.PaysArrive)
	datas["clientId"] = fmt.Sprint(commande.ClientId)

	return datas
}

func AddCommande(commande Commande) (int64, error) {
	return util.Create("Commande", getCommandeData(commande))
}

func UpdateCommande(commande Commande, id string) (int64, error) {
	return util.Update("Commande", getCommandeData(commande), fmt.Sprintf(" WHERE id = %v", id))
}

func GetAllCommandeByClient(clientId1 string) ([]Commande, error) {
	var id, etape, clientId, produitId int64
	var isImport, isExport bool
	var paysDepart, paysArrive, expediteurNom, expediteurAdresse, expediteurBP, expediteurTelephone, expediteurEmail, destinataireNom, destinataireAdresse, destinataireEmail, destinataireTelephone, destinataireBP, portChargement, paysChargement, portDechargement, paysDechargement, packingList, societe, NIF, date string
	variables := []interface{}{
		&id, &isImport, &isExport, &paysDepart, &paysArrive, &expediteurNom, &expediteurAdresse, &expediteurBP, &expediteurTelephone, &expediteurEmail, &destinataireNom, &destinataireAdresse, &destinataireEmail, &destinataireTelephone, &destinataireBP, &portChargement, &paysChargement, &portDechargement, &paysDechargement, &packingList, &societe, &NIF, &produitId, &etape, &clientId, &date,
	}
	keys := []string{"id", "isImport", "isExport", "paysDepart", "paysArrive", "expediteurNom", "expediteurAdresse", "expediteurBP", "expediteurTelephone", "expediteurEmail", "destinataireNom", "destinataireAdresse", "destinataireEmail", "destinataireTelephone", "destinataireBP", "portChargement", "paysChargement", "portDechargement", "paysDechargement", "packingList", "societe", "NIF", "produitId", "etape", "clientId", "date"}
	var commandes []Commande
	datas, err := util.ReadAll("Commande", variables, keys, fmt.Sprintf(` WHERE "clientId" = %v`, clientId1))
	if err != nil {
		return commandes, err
	}
	for _, data := range datas {
		var commande Commande
		commande.ID, _ = strconv.Atoi(fmt.Sprint((data["id"])))
		commande.ProduitId, _ = strconv.Atoi(fmt.Sprint((data["produitId"])))
		commande.IsImport, _ = strconv.ParseBool(fmt.Sprint(data["isImport"]))
		commande.IsExport, _ = strconv.ParseBool(fmt.Sprint(data["isExport"]))

		commande.ExpediteurAdresse = fmt.Sprint(data["expediteurAdresse"])
		commande.ExpediteurBP = fmt.Sprint(data["expediteurBP"])
		commande.ExpediteurEmail = fmt.Sprint(data["expediteurEmail"])
		commande.ExpediteurNom = fmt.Sprint(data["expediteurNom"])
		commande.ExpediteurTelephone = fmt.Sprint(data["expediteurTelephone"])

		commande.DestinataireAdresse = fmt.Sprint(data["destinataireAdresse"])
		commande.DestinataireBP = fmt.Sprint(data["destinataireBP"])
		commande.DestinataireEmail = fmt.Sprint(data["destinataireEmail"])
		commande.DestinataireNom = fmt.Sprint(data["destinataireNom"])
		commande.DestinataireTelephone = fmt.Sprint(data["destinataireTelephone"])

		commande.PortChargement = fmt.Sprint(data["portChargement"])
		commande.PaysChargement = fmt.Sprint(data["paysChargement"])
		commande.PortDechargement = fmt.Sprint(data["portDechargement"])
		commande.PaysDechargement = fmt.Sprint(data["paysDechargement"])
		commande.PackingList = fmt.Sprint(data["packingList"])

		commande.Societe = fmt.Sprint(data["societe"])
		commande.NIF = fmt.Sprint(data["NIF"])
		commande.Etape, _ = strconv.Atoi(fmt.Sprint(data["etape"]))
		commande.PaysDepart = fmt.Sprint(data["paysDepart"])
		commande.PaysArrive = fmt.Sprint(data["paysArrive"])
		commande.ClientId, _ = strconv.Atoi(fmt.Sprint(data["clientId"]))
		commande.Date = fmt.Sprint(data["date"])

		commande.Produit, _ = produit.GetProduitByID(fmt.Sprint(commande.ProduitId))

		commandes = append(commandes, commande)
	}
	return commandes, err
}

func GetAllCommande() ([]Commande, error) {
	var id, etape, clientId, produitId int64
	var isImport, isExport bool
	var paysDepart, paysArrive, expediteurNom, expediteurAdresse, expediteurBP, expediteurTelephone, expediteurEmail, destinataireNom, destinataireAdresse, destinataireEmail, destinataireTelephone, destinataireBP, portChargement, paysChargement, portDechargement, paysDechargement, packingList, societe, NIF, date string
	variables := []interface{}{
		&id, &isImport, &isExport, &paysDepart, &paysArrive, &expediteurNom, &expediteurAdresse, &expediteurBP, &expediteurTelephone, &expediteurEmail, &destinataireNom, &destinataireAdresse, &destinataireEmail, &destinataireTelephone, &destinataireBP, &portChargement, &paysChargement, &portDechargement, &paysDechargement, &packingList, &societe, &NIF, &produitId, &etape, &clientId, &date,
	}
	keys := []string{"id", "isImport", "isExport", "paysDepart", "paysArrive", "expediteurNom", "expediteurAdresse", "expediteurBP", "expediteurTelephone", "expediteurEmail", "destinataireNom", "destinataireAdresse", "destinataireEmail", "destinataireTelephone", "destinataireBP", "portChargement", "paysChargement", "portDechargement", "paysDechargement", "packingList", "societe", "NIF", "produitId", "etape", "clientId", "date"}
	var commandes []Commande
	datas, err := util.ReadAll("Commande", variables, keys, "")
	if err != nil {
		return commandes, err
	}
	for _, data := range datas {
		var commande Commande
		commande.ID, _ = strconv.Atoi(fmt.Sprint((data["id"])))
		commande.ProduitId, _ = strconv.Atoi(fmt.Sprint((data["produitId"])))
		commande.IsImport, _ = strconv.ParseBool(fmt.Sprint(data["isImport"]))
		commande.IsExport, _ = strconv.ParseBool(fmt.Sprint(data["isExport"]))

		commande.ExpediteurAdresse = fmt.Sprint(data["expediteurAdresse"])
		commande.ExpediteurBP = fmt.Sprint(data["expediteurBP"])
		commande.ExpediteurEmail = fmt.Sprint(data["expediteurEmail"])
		commande.ExpediteurNom = fmt.Sprint(data["expediteurNom"])
		commande.ExpediteurTelephone = fmt.Sprint(data["expediteurTelephone"])

		commande.DestinataireAdresse = fmt.Sprint(data["destinataireAdresse"])
		commande.DestinataireBP = fmt.Sprint(data["destinataireBP"])
		commande.DestinataireEmail = fmt.Sprint(data["destinataireEmail"])
		commande.DestinataireNom = fmt.Sprint(data["destinataireNom"])
		commande.DestinataireTelephone = fmt.Sprint(data["destinataireTelephone"])

		commande.PortChargement = fmt.Sprint(data["portChargement"])
		commande.PaysChargement = fmt.Sprint(data["paysChargement"])
		commande.PortDechargement = fmt.Sprint(data["portDechargement"])
		commande.PaysDechargement = fmt.Sprint(data["paysDechargement"])
		commande.PackingList = fmt.Sprint(data["packingList"])

		commande.Societe = fmt.Sprint(data["societe"])
		commande.NIF = fmt.Sprint(data["NIF"])
		commande.Etape, _ = strconv.Atoi(fmt.Sprint(data["etape"]))
		commande.PaysDepart = fmt.Sprint(data["paysDepart"])
		commande.PaysArrive = fmt.Sprint(data["paysArrive"])
		commande.ClientId, _ = strconv.Atoi(fmt.Sprint(data["clientId"]))
		commande.Date = fmt.Sprint(data["date"])

		commande.Produit, _ = produit.GetProduitByID(fmt.Sprint(commande.ProduitId))

		commandes = append(commandes, commande)
	}
	return commandes, err
}

func GetCommandeByID(paramId string) (Commande, error) {
	var id, etape, clientId, produitId int64
	var isImport, isExport bool
	var paysDepart, paysArrive, expediteurNom, expediteurAdresse, expediteurBP, expediteurTelephone, expediteurEmail, destinataireNom, destinataireAdresse, destinataireEmail, destinataireTelephone, destinataireBP, portChargement, paysChargement, portDechargement, paysDechargement, packingList, societe, NIF, date string
	variables := []interface{}{
		&id, &isImport, &isExport, &paysDepart, &paysArrive, &expediteurNom, &expediteurAdresse, &expediteurBP, &expediteurTelephone, &expediteurEmail, &destinataireNom, &destinataireAdresse, &destinataireEmail, &destinataireTelephone, &destinataireBP, &portChargement, &paysChargement, &portDechargement, &paysDechargement, &packingList, &societe, &NIF, &produitId, &etape, &clientId, &date,
	}
	keys := []string{"id", "isImport", "isExport", "paysDepart", "paysArrive", "expediteurNom", "expediteurAdresse", "expediteurBP", "expediteurTelephone", "expediteurEmail", "destinataireNom", "destinataireAdresse", "destinataireEmail", "destinataireTelephone", "destinataireBP", "portChargement", "paysChargement", "portDechargement", "paysDechargement", "packingList", "societe", "NIF", "produitId", "etape", "clientId", "date"}
	var commande Commande
	data, err := util.ReadOne("Commande", variables, keys, fmt.Sprintf(" WHERE id = %v", paramId))
	if err != nil {
		return commande, err
	}
	commande.ID, _ = strconv.Atoi(fmt.Sprint((data["id"])))
	commande.ProduitId, _ = strconv.Atoi(fmt.Sprint((data["produitId"])))
	commande.IsImport, _ = strconv.ParseBool(fmt.Sprint(data["isImport"]))
	commande.IsExport, _ = strconv.ParseBool(fmt.Sprint(data["isExport"]))

	commande.ExpediteurAdresse = fmt.Sprint(data["expediteurAdresse"])
	commande.ExpediteurBP = fmt.Sprint(data["expediteurBP"])
	commande.ExpediteurEmail = fmt.Sprint(data["expediteurEmail"])
	commande.ExpediteurNom = fmt.Sprint(data["expediteurNom"])
	commande.ExpediteurTelephone = fmt.Sprint(data["expediteurTelephone"])

	commande.DestinataireAdresse = fmt.Sprint(data["destinataireAdresse"])
	commande.DestinataireBP = fmt.Sprint(data["destinataireBP"])
	commande.DestinataireEmail = fmt.Sprint(data["destinataireEmail"])
	commande.DestinataireNom = fmt.Sprint(data["destinataireNom"])
	commande.DestinataireTelephone = fmt.Sprint(data["destinataireTelephone"])

	commande.PortChargement = fmt.Sprint(data["portChargement"])
	commande.PaysChargement = fmt.Sprint(data["paysChargement"])
	commande.PortDechargement = fmt.Sprint(data["portDechargement"])
	commande.PaysDechargement = fmt.Sprint(data["paysDechargement"])
	commande.PackingList = fmt.Sprint(data["packingList"])

	commande.Societe = fmt.Sprint(data["societe"])
	commande.NIF = fmt.Sprint(data["NIF"])
	commande.Etape, _ = strconv.Atoi(fmt.Sprint(data["etape"]))
	commande.PaysDepart = fmt.Sprint(data["paysDepart"])
	commande.PaysArrive = fmt.Sprint(data["paysArrive"])
	commande.ClientId, _ = strconv.Atoi(fmt.Sprint(data["clientId"]))
	commande.Date = fmt.Sprint(data["date"])

	commande.Produit, _ = produit.GetProduitByID(fmt.Sprint(commande.ProduitId))

	return commande, err
}

/** Colis */

func getColisData(colis Colis) map[string]string {
	datas := make(map[string]string)
	datas["largeur"] = fmt.Sprint(colis.Largeur)
	datas["hauteur"] = fmt.Sprint(colis.Hauteur)
	datas["longueur"] = fmt.Sprint(colis.Longueur)
	datas["poids"] = fmt.Sprint(colis.Poids)
	datas["photoURL"] = fmt.Sprint(colis.PhotoURL)
	datas["commandeId"] = fmt.Sprint(colis.CommandeId)

	return datas
}

func AddColis(colis Colis) (int64, error) {
	return util.Create("Colis", getColisData(colis))
}

func UpdateColis(colis Colis, id string) (int64, error) {
	return util.Update("Colis", getColisData(colis), fmt.Sprintf(" WHERE id = %v", id))
}

func GetAllByCommande(commandeId1 string) ([]Colis, error) {
	var id, quantite, commandeId int
	var largeur, longueur, hauteur, poids int
	var photoURL string

	variables := []interface{}{
		&id, &quantite, &largeur, &longueur, &hauteur, &poids, &commandeId, &photoURL,
	}
	keys := []string{
		"id", "quantite", "largeur", "longueur", "hauteur", "poids", "commandeId", "photoURL",
	}
	var colis []Colis
	datas, err := util.ReadAll("Colis", variables, keys, fmt.Sprintf(" WHERE commandeId = %v", commandeId))
	if err != nil {
		return colis, err
	}
	for _, data := range datas {
		var commande Colis
		commande.ID, _ = strconv.Atoi(fmt.Sprint((data["id"])))
		commande.Largeur, _ = strconv.Atoi(fmt.Sprint((data["largeur"])))
		commande.Longueur, _ = strconv.Atoi(fmt.Sprint(data["longueur"]))
		commande.Hauteur, _ = strconv.Atoi(fmt.Sprint(data["hauteur"]))
		commande.Poids, _ = strconv.Atoi(fmt.Sprint(data["poids"]))
		commande.Quantite, _ = strconv.Atoi(fmt.Sprint(data["quantite"]))
		commande.PhotoURL = fmt.Sprint(data["photoURL"])
		commande.CommandeId, _ = strconv.Atoi(fmt.Sprint((data["commandeId"])))

		colis = append(colis, commande)
	}
	return colis, err
}

/** Courrier */

func getCourrierData(courrier Courrier) map[string]string {
	datas := make(map[string]string)
	datas["largeur"] = fmt.Sprint(courrier.Largeur)
	datas["hauteur"] = fmt.Sprint(courrier.Hauteur)
	datas["longueur"] = fmt.Sprint(courrier.Longueur)
	datas["poids"] = fmt.Sprint(courrier.Poids)
	datas["photoURL"] = fmt.Sprint(courrier.PhotoURL)
	datas["commandeId"] = fmt.Sprint(courrier.CommandeId)
	datas["nom"] = fmt.Sprint(courrier.Nom)
	datas["type"] = fmt.Sprint(courrier.Type)

	return datas
}

func AddCourrier(courrier Courrier) (int64, error) {
	return util.Create("Courrier", getCourrierData(courrier))
}

func UpdateCourrier(courrier Courrier, id string) (int64, error) {
	return util.Update("Courrier", getCourrierData(courrier), fmt.Sprintf(" WHERE id = %v", id))
}

func GetAllCourrierByCommande(commandeId1 string) ([]Courrier, error) {
	var id, quantite, commandeId int
	var largeur, longueur, hauteur, poids int
	var photoURL, type1, nom string

	variables := []interface{}{
		&id, &quantite, &largeur, &longueur, &hauteur, &poids, &commandeId, &photoURL, &type1, &nom,
	}

	keys := []string{
		"id", "quantite", "largeur", "longueur", "hauteur", "poids", "commandeId", "photoURL", "type", "nom",
	}

	var courrier []Courrier
	datas, err := util.ReadAll("Courrier", variables, keys, fmt.Sprintf(" WHERE commandeId = %v", commandeId))
	if err != nil {
		return courrier, err
	}
	for _, data := range datas {
		var commande Courrier
		commande.ID, _ = strconv.Atoi(fmt.Sprint((data["id"])))
		commande.Largeur, _ = strconv.Atoi(fmt.Sprint((data["largeur"])))
		commande.Longueur, _ = strconv.Atoi(fmt.Sprint(data["longueur"]))
		commande.Hauteur, _ = strconv.Atoi(fmt.Sprint(data["hauteur"]))
		commande.Poids, _ = strconv.Atoi(fmt.Sprint(data["poids"]))
		commande.Quantite, _ = strconv.Atoi(fmt.Sprint(data["quantite"]))
		commande.PhotoURL = fmt.Sprint(data["photoURL"])
		commande.Nom = fmt.Sprint(data["nom"])
		commande.Type = fmt.Sprint(data["type"])
		commande.CommandeId, _ = strconv.Atoi(fmt.Sprint((data["commandeId"])))

		courrier = append(courrier, commande)
	}
	return courrier, err
}

/** Conteneur */

func getConteneurData(conteneur Conteneur) map[string]string {
	datas := make(map[string]string)
	datas["type"] = fmt.Sprint(conteneur.Type)
	datas["commandeId"] = fmt.Sprint(conteneur.CommandeId)

	return datas
}

func AddConteneur(conteneur Conteneur) (int64, error) {
	return util.Create("Conteneur", getConteneurData(conteneur))
}

func UpdateConteneur(conteneur Conteneur, id string) (int64, error) {
	return util.Update("Conteneur", getConteneurData(conteneur), fmt.Sprintf(" WHERE id = %v", id))
}

func GetAllConteneurByCommande(commandeId1 string) ([]Conteneur, error) {
	var id, commandeId int
	var type1 string

	variables := []interface{}{
		&id, &commandeId, &type1,
	}

	keys := []string{
		"id", "commandeId", "type",
	}

	var courrier []Conteneur
	datas, err := util.ReadAll("Conteneur", variables, keys, fmt.Sprintf(" WHERE commandeId = %v", commandeId))
	if err != nil {
		return courrier, err
	}
	for _, data := range datas {
		var commande Conteneur
		commande.ID, _ = strconv.Atoi(fmt.Sprint((data["id"])))
		commande.CommandeId, _ = strconv.Atoi(fmt.Sprint((data["commandeId"])))
		commande.Type = fmt.Sprint(data["type"])

		courrier = append(courrier, commande)
	}
	return courrier, err
}
