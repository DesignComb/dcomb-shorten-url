package model

func GetAllUrlShorten() ([]UrlShorten, error) {
	var urlShorten []UrlShorten

	tx := db.Find(&urlShorten)
	if tx.Error != nil {
		return []UrlShorten{}, tx.Error
	}

	return urlShorten, nil
}

func GetUrlShorten(id uint64) (UrlShorten, error) {
	var urlShorten UrlShorten

	tx := db.Where("id = ?", id).First(&urlShorten)

	if tx.Error != nil {
		return UrlShorten{}, tx.Error
	}

	return urlShorten, nil
}

func GetUrlShortenFromOrigin(origin string, userId uint64) []UrlShorten {
	var urlShorten []UrlShorten
	db.Where("origin = ?", origin).Where(db.Where("user_id is NULL").Or("user_id = ?", userId)).Find(&urlShorten)
	return urlShorten
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