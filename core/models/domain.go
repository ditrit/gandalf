package models

import (
	"github.com/jinzhu/gorm"
)

type Domain struct {
	gorm.Model
	Name string
}

func GetDescendants(database gorm.DB, id int) (domains []Domain) {
	//SELECT C.* FORM Comments C
	//JOIN TreePaths t
	//ON (c.commment_id = t.descendant)
	//WHERE t.ancestor = 4;
	database.Joins("JOIN domain_closure ON domain.id = domain_closure.descendant_id").Where("domain_closure.ancestor_id = ?", id).Find(&domains)

	return

}

func GetAncestors(database gorm.DB, id int) (domains []Domain) {
	//SELECT c.* FROM Comments c
	//JOIN TreePaths t
	//ON (c.commment_id = t.ancestor)
	//WHERE t.descendant = 6;
	database.Joins("JOIN domain_closure ON domain.id = domain_closure.ancestor_id").Where("domain_closure.descendant_id = ?", id).Find(&domains)

	return
}

func InsertNewChild(database gorm.DB, domain Domain, id int) {
	database.Save(domain)
	database.Where("name = ?", domain.Name).First(&domain)

	//INSERT INTO TreePaths (ancestor, descendant)
	//SELECT ancestor, 8 FROM TreePaths
	//WHERE descendant = 5
	//UNION ALL SELECT 8,8
	database.Raw("INSERT INTO domain_closure (ancestor_id, descendant_id) SELECT ancestor_id, ? FROM domain_closure WHERE descendant_id = ? UNION ALL SELECT ?,?", domain.ID, id, domain.ID, domain.ID)
}

func DeleteChild(database gorm.DB, domain Domain) {

	database.Delete(domain)
	//DELETE FORM TreePaths
	//WHERE descendant = 7
	var domainClosure DomainClosure
	database.Where("descendant = ?", domain.ID).Delete(&domainClosure)

}

func DeleteSubtree(database gorm.DB, id int) {

	database.Raw("DELETE FROM domain WHERE domain.id IN (SELECT descendant_id FROM domain_closure WHERE ancestor_id = ?)", id)
	//DELETE FROM TreePaths
	//WHERE descendant IN
	//(SELECT descendant FROM TreePaths
	//WHERE ancestor = 4);
	database.Raw("DELETE FROM domain_closure WHERE descendant_id IN (SELECT descendant_id FROM domain_closure WHERE ancestor_id = ?)", id)
	//database.Where("descendant in (?)", database.Table("domain_closure").Select("descendant").Where("ancestor = ?", id)).Delete()
}

//TODO ADD UPDATE NODE utile ?
func UpdateChild(database gorm.DB, domain Domain, id int) {

}

//TODO ADD UPDATE Subtree utile ?
func UpdateSubtree(database gorm.DB, domain Domain, id int) {

}
