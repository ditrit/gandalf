package models

import (
	"github.com/jinzhu/gorm"
)

type Domain struct {
	gorm.Model
	Name string
}

//TODO REVOIR ADD .ERROR + TRANSACTION
func GetDomainDescendants(database *gorm.DB, id uint) (domains []Domain) {
	database.Order("depth asc").Joins("JOIN domain_closures ON domains.id = domain_closures.descendant_id").Where("domain_closures.ancestor_id = ?", id).Find(&domains)

	return

}

func GetDomainAncestors(database *gorm.DB, id uint) (domains []Domain) {
	database.Order("depth desc").Joins("JOIN domain_closures ON domains.id = domain_closures.ancestor_id").Where("domain_closures.descendant_id = ?", id).Find(&domains)

	return
}

func InsertDomainRoot(database *gorm.DB, domain *Domain) {
	database.Save(domain)
	database.Where("name = ?", domain.Name).First(&domain)

	database.Exec("INSERT INTO domain_closures (ancestor_id, descendant_id, depth) SELECT ?,?,0;", domain.ID, domain.ID)
}

func InsertDomainNewChild(database *gorm.DB, domain *Domain, id uint) {
	database.Save(domain)
	database.Where("name = ?", domain.Name).First(&domain)

	database.Exec("INSERT INTO domain_closures (ancestor_id, descendant_id, depth) SELECT ancestor_id, ?, depth+1 FROM domain_closures WHERE descendant_id = ? UNION ALL SELECT ?,?,0;", domain.ID, id, domain.ID, domain.ID)
}

func DeleteDomainChild(database *gorm.DB, domain *Domain) {
	database.Delete(domain)

	var domainClosure DomainClosure
	database.Where("descendant = ?", domain.ID).Delete(&domainClosure)

}

func DeleteDomainSubtree(database *gorm.DB, id uint) {
	database.Exec("DELETE FROM domains WHERE domains.id IN (SELECT descendant_id FROM domain_closures WHERE ancestor_id = ?);", id)

	database.Exec("DELETE FROM domain_closures WHERE descendant_id IN (SELECT descendant_id FROM domain_closures WHERE ancestor_id = ?);", id)
}

func UpdateDomainChild(database *gorm.DB, domain *Domain, newAncestor uint) {
	//database.Exec("DELETE a.* FROM domain_closures AS a JOIN domain_closures AS d ON a.descendant_id = d.descendant_id LEFT JOIN domain_closures AS x ON x.ancestor_id = d.ancestor_id AND x.descendant_id = a.ancestor_id WHERE d.ancestor_id = ? AND x.ancestor_id IS NULL;", oldAncestor)
	//database.Exec("DELETE FROM domain_closures WHERE EXISTS (SELECT a.descendant_id, d.ancestor_id FROM domain_closures AS a JOIN domain_closures AS d ON a.descendant_id = d.descendant_id LEFT JOIN domain_closures AS x ON x.ancestor_id = d.ancestor_id AND x.descendant_id = a.ancestor_id WHERE d.ancestor_id = ? AND x.ancestor_id IS NULL);", oldAncestor)
	//database.Exec("INSERT INTO domain_closures (ancestor_id, descendant_id, depth) SELECT supertree.ancestor_id, subtree.descendant_id, supertree.depth + subtree.depth + 1 FROM domain_closures AS supertree JOIN domain_closures AS subtree WHERE subtree.ancestor_id = ? AND supertree.descendant_id = ?;", oldAncestor, newAncestor)
	//database.Exec("INSERT INTO domain_closures (ancestor_id, descendant_id, depth) SELECT supertree.ancestor_id, subtree.descendant_id, supertree.depth + subtree.depth + 1 FROM domain_closures AS supertree JOIN domain_closures AS subtree WHERE subtree.ancestor_id = ? AND supertree.descendant_id = ?;", oldAncestor, newAncestor)

	//database.Exec("DELETE FROM domain_closures WHERE descendant_id IN (SELECT descendant_id FROM domain_closures WHERE ancestor_id = ? AND descendant_id <> ?);", oldAncestor, oldAncestor)
	database.Exec("DELETE FROM domain_closures WHERE descendant_id IN (SELECT descendant_id FROM domain_closures WHERE ancestor_id = ?) AND ancestor_id IN (SELECT ancestor_id FROM domain_closures WHERE descendant_id = ? AND ancestor_id != descendant_id);", domain.ID, domain.ID)

	database.Exec("INSERT INTO domain_closures (ancestor_id, descendant_id, depth) SELECT supertree.ancestor_id, subtree.descendant_id,  supertree.depth + subtree.depth + 1 AS depth FROM domain_closures AS supertree CROSS JOIN domain_closures AS subtree WHERE supertree.descendant_id = ? AND subtree.ancestor_id = ?;", newAncestor, domain.ID)
}
