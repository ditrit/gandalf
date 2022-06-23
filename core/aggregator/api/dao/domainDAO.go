package dao

import (
	"errors"
	"log"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/google/uuid"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListDomain(database *gorm.DB) (domains []models.Domain, err error) {
	err = database.Preload("Parent").Preload("Products").Preload("Libraries").Preload("Authorizations.User").Preload("Authorizations.Role").Preload("Tags").Preload("Environments").Find(&domains).Error
	log.Println(err)
	return
}

func CreateDomain(database *gorm.DB, domain *models.Domain, parentDomainID uuid.UUID) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			domain.ParentID = parentDomainID
			err = database.Save(&domain).Error
		} else {
			err = errors.New("invalid state")
		}
	}

	return
}

func TreeDomain(database *gorm.DB) (result *models.Domain, err error) {
	var results []*models.Domain
	database.Raw("select * from domains order by parent_id").Scan(&results)

	domain := results[0]
	TreeRecursiveDomain(domain, results)

	result = domain
	return
}

func TreeRecursiveDomain(domain *models.Domain, results []*models.Domain) {
	for _, result := range results {
		if result.ParentID == domain.ID {
			currentDomain := result
			domain.Childs = append(domain.Childs, currentDomain)
		}
	}
	for _, child := range domain.Childs {
		TreeRecursiveDomain(child, results)
	}
}

func ReadDomain(database *gorm.DB, id uuid.UUID) (domain models.Domain, err error) {
	err = database.Where("id = ?", id).Preload("Parent").Preload("Products").Preload("Libraries").Preload("Authorizations.User").Preload("Authorizations.Role").Preload("Tags").Preload("Environments.EnvironmentType").Preload("Environments.Domain").First(&domain).Error

	return
}

func ReadDomainByName(database *gorm.DB, name string) (domain models.Domain, err error) {
	log.Println("DAO")
	err = database.Preload("Parent").Preload("Products").Preload("Libraries").Preload("Authorizations.User").Preload("Authorizations.Role").Preload("Tags").Preload("Environments.EnvironmentType").Preload("Environments.Domain").Where("name = ?", name).First(&domain).Error
	log.Println(err)
	return
}

func UpdateDomain(database *gorm.DB, domain models.Domain) (err error) {
	err = database.Save(&domain).Error

	return
}

func DeleteDomain(database *gorm.DB, id uuid.UUID) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var domain models.Domain
			err = database.Where("id = ?", id).First(&domain).Error
			if err == nil {
				err = database.Unscoped().Delete(&domain).Error
			}
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ListDomainTag(database *gorm.DB, domain models.Domain) (tags []models.Tag, err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Model(&domain).Association("Tags").Find(&tags).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func AddDomainTag(database *gorm.DB, domain models.Domain, tag models.Tag) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Model(&domain).Association("Tags").Append(&tag).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func RemoveDomainTag(database *gorm.DB, domain models.Domain, tag models.Tag) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Model(&domain).Association("Tags").Delete(&tag).Error

		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ListDomainLibrary(database *gorm.DB, domain models.Domain) (libraries []models.Library, err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Model(&domain).Association("Libraries").Find(&libraries).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func AddDomainLibrary(database *gorm.DB, domain models.Domain, library models.Library) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Model(&domain).Association("Libraries").Append(&library).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func RemoveDomainLibrary(database *gorm.DB, domain models.Domain, library models.Library) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Model(&domain).Association("Libraries").Delete(&library).Error

		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
