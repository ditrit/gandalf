package dao

import (
	"github.com/ditrit/gandalf/core/cluster/api/utils"

	"github.com/ditrit/gandalf/core/models"

	"github.com/jinzhu/gorm"
)

func ListAdminTenant(database *gorm.DB) (users []models.User, err error) {

	var authorizations []models.Authorization

	var admin models.Role
	if err := database.Where("name = ?", "Administrator").First(&admin).Error; err != nil {
		// return any error will rollback
		var root models.Domain
		if err := database.Where("name = ?", "Root").First(&root).Error; err != nil {
			// return any error will rollback
			if err := database.Where("role_id = ? and domain_id = ?", admin.ID, root.ID).Preload("User").Preload("Role").Preload("Domain").Find(&authorizations).Error; err != nil {
				// return any error will rollback
				for _, authorization := range authorizations {
					users = append(users, authorization.User)
				}
			}
		}
	}

	return
}

func CreateAdminTenant(database *gorm.DB, user models.User) (err error) {

	err = database.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&user).Error; err != nil {
			// return any error will rollback
			return err
		}
		var admin models.Role
		if err := tx.Where("name = ?", "Administrator").First(&admin).Error; err != nil {
			// return any error will rollback
			return err
		}
		var root models.Domain
		if err := tx.Where("name = ?", "Root").First(&root).Error; err != nil {
			// return any error will rollback
			return err
		}
		authorization := models.Authorization{User: user, Role: admin, Domain: root}
		if err := tx.Create(&authorization).Error; err != nil {
			// return any error will rollback
			return err
		}

		return nil
	})

	if err == nil {
		err = utils.ChangeStateGandalf(database)
	}

	return
}

/* func ReadAdminTenant(database *gorm.DB, id int) (user models.User, err error) {
	err = database.First(&user, id).Error

	return
}

func UpdateAdminTenant(database *gorm.DB, user models.User) (err error) {
	err = database.Save(&user).Error

	return
}

func DeleteAdminTenant(database *gorm.DB, id int) (err error) {
	var user models.User
	err = database.Delete(&user, id).Error

	return
}
*/
