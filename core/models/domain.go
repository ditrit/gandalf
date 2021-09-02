package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Domain struct {
	gorm.Model
	Name string `gorm:"not null"`
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
	err = database.Transaction(func(tx *gorm.DB) error {

		if err := tx.Save(&domain).Error; err != nil {
			// return any error will rollback
			return err
		}
		if err := tx.Where("name = ?", domain.Name).First(&domain).Error; err != nil {
			// return any error will rollback
			return err
		}
		if err := tx.Exec("INSERT INTO domain_closures (ancestor_id, descendant_id, depth) SELECT ?,?,0;", domain.ID, domain.ID).Error; err != nil {
			// return any error will rollback
			return err
		}
		return nil
	})

	return
}

func InsertDomainNewChild(database *gorm.DB, domain Domain, id uint) (err error) {

	err = database.Transaction(func(tx *gorm.DB) error {

		if err := tx.Save(&domain).Error; err != nil {
			// return any error will rollback
			return err
		}
		if err := tx.Where("name = ?", domain.Name).First(&domain).Error; err != nil {
			// return any error will rollback
			return err
		}
		if err := tx.Exec("INSERT INTO domain_closures (ancestor_id, descendant_id, depth) SELECT ancestor_id, ?, depth+1 FROM domain_closures WHERE descendant_id = ? UNION ALL SELECT ?,?,0;", domain.ID, id, domain.ID, domain.ID).Error; err != nil {
			// return any error will rollback
			return err
		}
		return nil
	})

	return
}

func DeleteDomainChild(database *gorm.DB, domain Domain) (err error) {

	err = database.Transaction(func(tx *gorm.DB) error {

		if err := tx.Delete(&Domain{}, int(domain.ID)).Error; err != nil {
			fmt.Println(err)
			// return any error will rollback
			return err
		}
		var domainClosure DomainClosure
		if err := tx.Where("descendant_id = ?", domain.ID).Delete(&domainClosure).Error; err != nil {
			// return any error will rollback
			fmt.Println(err)
			return err
		}
		return nil
	})
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
