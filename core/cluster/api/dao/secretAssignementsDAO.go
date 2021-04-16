package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/cluster/api/utils"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListSecretAssignement(database *gorm.DB) (secretAssignement []models.SecretAssignement, err error) {
	err = database.Find(&secretAssignement).Error

	return
}

func CreateSecretAssignement(database *gorm.DB, secretAssignement models.SecretAssignement) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Create(&secretAssignement).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func DeleteSecretAssignement(database *gorm.DB, secret string) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var secret models.SecretAssignement
			err = database.Delete(&secret, secret).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
