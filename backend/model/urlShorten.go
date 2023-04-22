package model

func GetAllUrlShorten() ([]UrlShorten, error) {
	var urlShorten []UrlShorten

	tx := db.Find(&urlShorten)
	if tx.Error != nil {
		return []UrlShorten{}, tx.Error
	}

	return urlShorten, nil
}

func GetUserUrlShorten(id uint64, userId uint64) (UrlShorten, error) {
	var urlShorten UrlShorten
	tx := db.Where("id = ? AND user_id = ?", id, userId).First(&urlShorten)

	if tx.Error != nil {
		return UrlShorten{}, tx.Error
	}

	return urlShorten, nil
}

func GetUserUrlShortenFromOrigin(origin string, userId uint64) []UrlShorten {
	var userUrlShorten []UrlShorten
	db.Where("origin = ?", origin).Where("user_id = ?", userId).Find(&userUrlShorten)
	return userUrlShorten
}

func GetNonUserUrlShortenFromOrigin(origin string) []UrlShorten {
	var nonUserUrlShorten []UrlShorten
	db.Where("origin = ?", origin).Where("user_id is NULL").Find(&nonUserUrlShorten)
	return nonUserUrlShorten
}

func CreateUrlShorten(urlShorten UrlShorten) (UrlShorten, error) {
	tx := db.Create(&urlShorten)
	return urlShorten, tx.Error
}

func UpdateUrlShorten(urlShorten UrlShorten) error {
	tx := db.Save(&urlShorten)
	return tx.Error
}

func DeleteUrlShorten(id uint64) error {
	tx := db.Unscoped().Delete(&UrlShorten{}, id)
	return tx.Error
}

func FindByUrl(url string) (UrlShorten, error) {
	var urlShorten UrlShorten
	tx := db.Where("short = ?", url).First(&urlShorten)
	return urlShorten, tx.Error
}

func SearchUserUrlShorten(userId uint64, keyword string) []UrlShorten {
	var urlShorten []UrlShorten
	db.Where("user_id = ? AND (origin LIKE ? OR short LIKE ?)", userId, "%"+keyword+"%", "%"+keyword+"%").Find(&urlShorten)
	return urlShorten
}

func SearchNonUserUrlShorten(keyword string) []UrlShorten {
	var urlShorten []UrlShorten
	db.Where("user_id is NULL AND (origin LIKE ? OR short LIKE ?)", "%"+keyword+"%", "%"+keyword+"%").Find(&urlShorten)
	return urlShorten
}
