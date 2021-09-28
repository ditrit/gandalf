package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Domain struct {
	gorm.Model
	ParentID         *uint
	Parent           *Domain `gorm:"constraint:OnDelete:CASCADE;"`
	Name             string  `gorm:"not null"`
	ShortDescription string
	Description      string
	Logo             string
}

func (d *Domain) BeforeDelete(tx *gorm.DB) (err error) {
	var childs []Domain
	tx.Where("parent_id = ?", d.ID).Find(&childs)
	fmt.Println(childs)
	for _, child := range childs {
		tx.Unscoped().Delete(&child)
	}

	return
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
