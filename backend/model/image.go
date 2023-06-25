package model

func CreateImage(image Image) (Image, error) {
	tx := db.Create(&image)
	return image, tx.Error
}

func FindImage(id uint64) (Image, error) {
	var image Image
	tx := db.First(&image, id)
	return image, tx.Error
}