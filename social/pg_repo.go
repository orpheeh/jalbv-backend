package social

import (
	"fmt"

	util "github.com/orpheeh/jalbv-backend/utils"
)

var tableName = "Social"

func GetAllSocial() ([]Social, error) {
	var id, libelle, imageURL, link, isActive string
	variables := []interface{}{
		&id, &libelle, &imageURL, &link, &isActive,
	}
	keys := []string{"id", "libelle", "imageURL", "link", "isActive"}
	var socials []Social
	datas, err := util.ReadAll(tableName, variables, keys, "")
	if err != nil {
		return socials, err
	}
	for _, data := range datas {
		var social Social
		social.ID = data["id"]
		social.Libelle = string(data["libelle"])
		social.ImageURL = string(data["imageURL"])
		social.Link = string(data["link"])
		social.IsActive = string(data["isActive"])
		socials = append(socials, social)
	}
	return socials, err
}

func UpdateSocial(social Social, id string) (int64, error) {
	datas := make(map[string]string)
	if social.ImageURL != "" {
		datas["imageURL"] = social.ImageURL
	}
	if social.IsActive != "" {
		datas["isActive"] = social.IsActive
	}
	if social.Libelle != "" {
		datas["libelle"] = social.Libelle
	}
	if social.Link != "" {
		datas["link"] = social.Link
	}
	return util.Update(tableName, datas, fmt.Sprintf(" WHERE id = %v", id))
}
