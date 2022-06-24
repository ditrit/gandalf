package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/google/uuid"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)


func ListRequirementGroups(database *gorm.DB) (requirementGroups []models.RequirementGroup, err error) {
	err = database.Find(&requirementGroups).Error

	return
}

func CreateRequirementGroup(database *gorm.DB, requirementGroup *models.RequirementGroup) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Create(&requirementGroup).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadRequirementGroup(database *gorm.DB, id uuid.UUID) (requirementGroup models.RequirementGroup, err error) {
	err = database.Where("id = ?", id).First(&requirementGroup).Error

	return
}

func UpdateRequirementGroup(database *gorm.DB, requirementGroup models.RequirementGroup) (err error) {
	err = database.Save(&requirementGroup).Error

	return
}

func DeleteRequirementGroup(database *gorm.DB, id uuid.UUID) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var requirementGroup models.RequirementGroup
			err = database.Where("id = ?", id).First(&requirementGroup).Error
			if err == nil {
				err = database.Unscoped().Delete(&requirementGroup).Error
			}
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
