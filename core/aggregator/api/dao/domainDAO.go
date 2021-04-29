package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListDomain(database *gorm.DB) (domains []models.Domain, err error) {
	var root models.Domain
	err = database.Where("name = ?", "root").First(&root).Error
	if err == nil {
		domains, err = models.GetDomainDescendants(database, root.ID)
	}

	return
}

func CreateDomain(database *gorm.DB, domain models.Domain, parentDomainName string) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			if parentDomainName == "root" {
				models.InsertDomainRoot(database, domain)
			} else {
				var parentDomain models.Domain
				err = database.Where("name = ?", parentDomainName).First(&parentDomain).Error
				if err == nil {
					models.InsertDomainNewChild(database, domain, parentDomain.ID)
				}
			}
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadDomain(database *gorm.DB, id int) (domain models.Domain, err error) {
	err = database.First(&domain, id).Error

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
				models.DeleteDomainChild(database, domain)
			}
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
