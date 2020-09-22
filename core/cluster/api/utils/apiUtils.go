package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ditrit/gandalf/core/models"

	"github.com/ditrit/gandalf/core/cluster/database"

	"github.com/jinzhu/gorm"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetDatabase(mapDatabase map[string]*gorm.DB, databasePath, tenant string) *gorm.DB {
	if _, ok := mapDatabase[tenant]; !ok {
		var databaseCreated, err = database.IsDatabaseCreated(databasePath, tenant)
		if err == nil {
			fmt.Println("databaseCreated")
			fmt.Println(databaseCreated)
			if databaseCreated {
				var tenantDatabaseClient *gorm.DB
				tenantDatabaseClient, err = database.NewTenantDatabaseClient(tenant, databasePath)
				if err == nil {
					mapDatabase[tenant] = tenantDatabaseClient
				} else {
					log.Println("Can't create database client")
				}
			} else {
				return nil
			}
		} else {
			log.Println("Can't detect if the database is created or not")
		}
	}

	return mapDatabase[tenant]
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func GetStateGandalf(client *gorm.DB) (bool, error) {
	var state models.State
	err := client.First(&state).Error
	if err == nil {
		return state.Admin, err
	}
	return false, err

}

func ChangeStateGandalf(client *gorm.DB) (err error) {
	var state models.State
	err = client.First(&state).Error
	if err == nil {
		if !state.Admin {
			var roleadmin models.Role
			err = client.Where("name = ?", "Administrator").First(&roleadmin).Error
			if err == nil {
				var users []models.User
				result := client.Where("role_id = ?", roleadmin.ID).Preload("Role").Find(&users)
				if result.Error == nil {
					if result.RowsAffected >= 2 {
						state.Admin = true
						client.Save(&state)
					}
				}
			}
		}
	}
	return err
}

//TODO REVOIR UTILE ???
/* func ChangeStateTenant(client *gorm.DB) {
	var state models.State
	client.First(&state)

	if !state.Admin {
		var roleadmin models.Role
		client.Where("name = ?", "Administrator").First(&roleadmin)
		var users []models.User
		result := client.Where("role_id = ?", roleadmin.ID).Find(&users)
		if result.RowsAffected >= 2 {
			state.Admin = true
			client.Update(&state)
		}
	}
} */
