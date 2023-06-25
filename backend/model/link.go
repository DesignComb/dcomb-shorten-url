package model

func GetTreeAllLink(treeId uint64) ([]Link, error) {
	var link []Link

	tx := db.Where("tree_id = ?", treeId).Find(&link)

	if tx.Error != nil {
		return []Link{}, tx.Error
	}

	return link, nil
}