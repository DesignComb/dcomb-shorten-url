package model

func GetTreeAllLink(treeId uint64) ([]Link, error) {
	var link []Link

	tx := db.Where("tree_id = ?", treeId).Find(&link)

	if tx.Error != nil {
		return []Link{}, tx.Error
	}

	return link, nil
}

func CreateLinks(links []Link) ([]Link, error) {
	tx := db.Create(&links)
	return links, tx.Error
}

func DeleteTreeLinks(treeId uint64) (error) {
	tx := db.Where("tree_id = ?", treeId).Delete(&Link{})
	return tx.Error
}