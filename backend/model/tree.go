package model

func FindTree(id uint64) (Tree, error) {
	var tree Tree
	tx := db.First(&tree, id)
	return tree, tx.Error
}

func CreateTree(tree Tree) (Tree, error) {
	tx := db.Create(&tree)
	return tree, tx.Error
}

func UpdateTree(tree Tree) error {
	tx := db.Save(&tree)
	return tx.Error
}