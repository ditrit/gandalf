package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Domain struct {
	gorm.Model
	ParentID         uint    `gorm:"constraint:OnDelete:CASCADE;"`
	Parent           *Domain `gorm:"constraint:OnDelete:CASCADE;"`
	Name             string  `gorm:"not null"`
	ShortDescription string
	Description      string
	Logo             string
}

func GetDomainTree(database *gorm.DB, id uint) (domains []Domain, err error) {

	var domainTree = new(DomainTree)
	domainTree.Name = "root"
	fmt.Println("toto2")
	//err = database.Order("depth asc").Joins("JOIN domain_closures ON domains.id = domain_closures.descendant_id").Where("domain_closures.ancestor_id = ?", id).Find(&domains).Error
	//rows, err := database.Table("domains").Select("domains.id, domains.name, domain_closures.ancestor_id , domain_closures.descendant_id , domain_closures.depth").Joins("JOIN domain_closures ON domains.id = domain_closures.descendant_id").Where("domain_closures.ancestor_id = ?", id).Order("depth asc").Rows()
	//rows, err := database.Raw("SELECT d.id, d.name, domain_closures.ancestor_id, domain_closures.depth FROM domain_closures JOIN domains as d ON (domain_closures.descendant_id = d.id) JOIN domains as t ON (domain_closures.ancestor_id = t.id) WHERE domain_closures.depth <= 1 ORDER BY domain_closures.ancestor_id;").Rows()
	rows, err := database.Raw("select domains.*, domain_closures.ancestor_id  from domains join domain_closures on domains.id = domain_closures.descendant_id WHERE domain_closures.depth = 1 or (domain_closures.depth = 0 and domains.name = 'root');").Rows()

	//rows, err := database.Raw("select n.name from domain_closures d join domain_closures a on (a.descendant_id = d.descendant_id) join domains n on (n.id = a.ancestor_id) where d.ancestor_id = ? and d.descendant_id != d.ancestor_id group by d.descendant_id, n.name;", id).Rows()
	fmt.Println("err")
	fmt.Println(err)
	fmt.Println("rows")
	for rows.Next() {
		fmt.Println(rows)
	}
	/* var currentDomainTree *DomainTree
	var rowname string
	for rows.Next() {

		fmt.Println(rows)
		rows.Scan(&rowname)
		if rowname == "root" {
			currentDomainTree = domainTree
		} else {
			//loopChilds := currentDomainTree.Childs
			find := false
			for i, child := range currentDomainTree.Childs {
				if rowname == child.Name {
					currentDomainTree = currentDomainTree.Childs[i]
					find = true
				}
			}
			if !find {
				createdomainTree := new(DomainTree)
				createdomainTree.Name = rowname
				currentDomainTree.Childs = append(currentDomainTree.Childs, createdomainTree)
			}
		}
	}

	fmt.Println(domainTree)
	for _, child := range domainTree.Childs {
		fmt.Println(child)
		fmt.Println(child.Childs)
	} */

	return

}

func GetDomainDescendants(database *gorm.DB, id uint) (domains []Domain, err error) {
	err = database.Order("depth asc").Joins("JOIN domain_closures ON domains.id = domain_closures.descendant_id").Where("domain_closures.ancestor_id = ?", id).Find(&domains).Error

	return

}

func GetDomainAncestors(database *gorm.DB, id uint) (domains []Domain, err error) {
	err = database.Order("depth desc").Joins("JOIN domain_closures ON domains.id = domain_closures.ancestor_id").Where("domain_closures.descendant_id = ?", id).Find(&domains).Error

	return
}

func InsertDomainRoot(database *gorm.DB, domain Domain) (err error) {
	err = database.Save(&domain).Error

	return
}

func InsertDomainNewChild(database *gorm.DB, domain Domain, id uint) (err error) {
	domain.ParentID = id
	err = database.Save(&domain).Error

	return
}

func DeleteDomainChild(database *gorm.DB, domain Domain) (err error) {
	err = database.Unscoped().Delete(&Domain{}, int(domain.ID)).Error
	return
}

func DeleteDomainSubtree(database *gorm.DB, id uint) (err error) {

	err = database.Transaction(func(tx *gorm.DB) error {

		if err := tx.Exec("DELETE FROM domains WHERE domains.id IN (SELECT descendant_id FROM domain_closures WHERE ancestor_id = ?);", id).Error; err != nil {
			// return any error will rollback
			return err
		}
		if err := tx.Exec("DELETE FROM domain_closures WHERE descendant_id IN (SELECT descendant_id FROM domain_closures WHERE ancestor_id = ?);", id).Error; err != nil {
			// return any error will rollback
			return err
		}
		return nil
	})

	return
}

func UpdateDomainChild(database *gorm.DB, domain Domain, newAncestor uint) (err error) {

	err = database.Transaction(func(tx *gorm.DB) error {

		if err := tx.Exec("DELETE FROM domain_closures WHERE descendant_id IN (SELECT descendant_id FROM domain_closures WHERE ancestor_id = ?) AND ancestor_id IN (SELECT ancestor_id FROM domain_closures WHERE descendant_id = ? AND ancestor_id != descendant_id);", domain.ID, domain.ID).Error; err != nil {
			// return any error will rollback
			return err
		}
		if err := tx.Exec("INSERT INTO domain_closures (ancestor_id, descendant_id, depth) SELECT supertree.ancestor_id, subtree.descendant_id,  supertree.depth + subtree.depth + 1 AS depth FROM domain_closures AS supertree CROSS JOIN domain_closures AS subtree WHERE supertree.descendant_id = ? AND subtree.ancestor_id = ?;", newAncestor, domain.ID).Error; err != nil {
			// return any error will rollback
			return err
		}
		return nil
	})

	return
}
