package model

func FindTree(id uint64) (Tree, error) {
	var tree Tree
	tx := db.First(&tree, id)
	return tree, tx.Error
}
