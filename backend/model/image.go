package model

func CreateImage(image Image) (Image, error) {
	tx := db.Create(&image)
	return image, tx.Error
}