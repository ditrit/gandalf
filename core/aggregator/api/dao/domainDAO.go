package dao

import (
	"errors"
	"fmt"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListDomain(database *gorm.DB) (domains []models.Domain, err error) {
	err = database.Preload("Parent").Preload("Products").Preload("Library").Find(&domains).Error
	fmt.Println(err)
	return
}

func CreateDomain(database *gorm.DB, domain models.Domain, parentDomainID uint) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			// if parentDomainName == "root" {
			// 	err = models.InsertDomainRoot(database, domain)
			// } else {
			// var parentDomain models.Domain
			// err = database.Where("name = ?", parentDomainName).First(&parentDomain).Error
			// if err == nil {
			domain.ParentID = parentDomainID
			err = database.Save(&domain).Error
			//}
			//}
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func TreeDomain(database *gorm.DB) (result *models.DomainTree, err error) {
	var results []models.Domain
	database.Raw("select * from domains order by parent_id").Scan(&results)

	domainTree := new(models.DomainTree)
	domainTree.Domain = results[0]
	TreeRecursiveDomain(domainTree, results)

	result = domainTree
	return
}

func TreeRecursiveDomain(domaintree *models.DomainTree, results []models.Domain) {
	for _, result := range results {
		if result.ParentID == domaintree.Domain.ID {
			currentDomainTree := new(models.DomainTree)
			currentDomainTree.Domain = result
			domaintree.Childs = append(domaintree.Childs, currentDomainTree)
		}
	}
	for _, child := range domaintree.Childs {
		TreeRecursiveDomain(child, results)
	}
}

func ReadDomain(database *gorm.DB, id int) (domain models.Domain, err error) {
	err = database.Preload("Parent").Preload("Products").Preload("Library").First(&domain, id).Error

	return
}

func ReadDomainByName(database *gorm.DB, name string) (domain models.Domain, err error) {
	fmt.Println("DAO")
	err = database.Preload("Parent").Preload("Products").Preload("Library").Where("name = ?", name).First(&domain).Error
	fmt.Println(err)
	fmt.Println(domain)
	return
}

func UpdateDomain(database *gorm.DB, domain models.Domain) (err error) {
	err = database.Save(&domain).Error

	return
}

func DeleteDomain(database *gorm.DB, id int) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var domain models.Domain
			err = database.First(&domain, id).Error
			if err == nil {
				err = database.Unscoped().Delete(&domain).Error
			}
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
