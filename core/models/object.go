package models

import "github.com/jinzhu/gorm"

type Object struct {
	gorm.Model
	Name    string
	Schema  string
	Actions []Action `gorm:"many2many:object_actions;"`
	Domains []Domain `gorm:"many2many:object_domains;"`
}

func GetObjectDescendants(database *gorm.DB, id uint) (objects []Object) {
	database.Order("depth asc").Joins("JOIN object_closures ON objects.id = object_closures.descendant_id").Where("object_closures.ancestor_id = ?", id).Preload("Domains").Preload("Actions").Find(&objects)

	return

}

func GetObjectAncestors(database *gorm.DB, id uint) (objects []Object) {
	database.Order("depth desc").Joins("JOIN object_closures ON objects.id = object_closures.ancestor_id").Where("object_closures.descendant_id = ?", id).Preload("Domains").Preload("Actions").Find(&objects)

	return
}

func InsertObjectRoot(database *gorm.DB, object *Object) {
	database.Save(object)
	database.Where("name = ?", object.Name).First(&object)

	database.Exec("INSERT INTO object_closures (ancestor_id, descendant_id, depth) SELECT ?,?,0;", object.ID, object.ID)
}

func InsertObjectNewChild(database *gorm.DB, object *Object, id uint) {
	database.Save(object)
	database.Where("name = ?", object.Name).First(&object)

	database.Exec("INSERT INTO object_closures (ancestor_id, descendant_id, depth) SELECT ancestor_id, ?, depth+1 FROM object_closures WHERE descendant_id = ? UNION ALL SELECT ?,?,0;", object.ID, id, object.ID, object.ID)
}

func DeleteObjectChild(database *gorm.DB, object *Object) {
	database.Delete(object)

	var objectClosure ObjectClosure
	database.Where("descendant = ?", object.ID).Delete(&objectClosure)

}

func DeleteObjectSubtree(database *gorm.DB, id uint) {
	database.Exec("DELETE FROM objects WHERE objects.id IN (SELECT descendant_id FROM object_closures WHERE ancestor_id = ?);", id)

	database.Exec("DELETE FROM object_closures WHERE descendant_id IN (SELECT descendant_id FROM object_closures WHERE ancestor_id = ?);", id)
}

func UpdateObjectChild(database *gorm.DB, object *Object, newAncestor uint) {
	//database.Exec("DELETE a.* FROM object_closures AS a JOIN object_closures AS d ON a.descendant_id = d.descendant_id LEFT JOIN object_closures AS x ON x.ancestor_id = d.ancestor_id AND x.descendant_id = a.ancestor_id WHERE d.ancestor_id = ? AND x.ancestor_id IS NULL;", oldAncestor)
	//database.Exec("DELETE FROM object_closures WHERE EXISTS (SELECT a.descendant_id, d.ancestor_id FROM object_closures AS a JOIN object_closures AS d ON a.descendant_id = d.descendant_id LEFT JOIN object_closures AS x ON x.ancestor_id = d.ancestor_id AND x.descendant_id = a.ancestor_id WHERE d.ancestor_id = ? AND x.ancestor_id IS NULL);", oldAncestor)
	//database.Exec("INSERT INTO object_closures (ancestor_id, descendant_id, depth) SELECT supertree.ancestor_id, subtree.descendant_id, supertree.depth + subtree.depth + 1 FROM object_closures AS supertree JOIN object_closures AS subtree WHERE subtree.ancestor_id = ? AND supertree.descendant_id = ?;", oldAncestor, newAncestor)
	//database.Exec("INSERT INTO object_closures (ancestor_id, descendant_id, depth) SELECT supertree.ancestor_id, subtree.descendant_id, supertree.depth + subtree.depth + 1 FROM object_closures AS supertree JOIN object_closures AS subtree WHERE subtree.ancestor_id = ? AND supertree.descendant_id = ?;", oldAncestor, newAncestor)

	//database.Exec("DELETE FROM object_closures WHERE descendant_id IN (SELECT descendant_id FROM object_closures WHERE ancestor_id = ? AND descendant_id <> ?);", oldAncestor, oldAncestor)
	database.Exec("DELETE FROM object_closures WHERE descendant_id IN (SELECT descendant_id FROM object_closures WHERE ancestor_id = ?) AND ancestor_id IN (SELECT ancestor_id FROM object_closures WHERE descendant_id = ? AND ancestor_id != descendant_id);", object.ID, object.ID)

	database.Exec("INSERT INTO object_closures (ancestor_id, descendant_id, depth) SELECT supertree.ancestor_id, subtree.descendant_id,  supertree.depth + subtree.depth + 1 AS depth FROM object_closures AS supertree CROSS JOIN object_closures AS subtree WHERE supertree.descendant_id = ? AND subtree.ancestor_id = ?;", newAncestor, object.ID)
}
