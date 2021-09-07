package utils

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/ditrit/gandalf/core/models"

	"github.com/jinzhu/gorm"
)

var jwtKey = []byte("secret_key")

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetJwtKey() []byte {
	return jwtKey
}

/* //TODO
// GetDatabaseClientByTenant : Cluster database client getter by tenant.
func GetDatabaseClientByTenant(tenant, addr string, mapDatabaseClient map[string]*gorm.DB) *gorm.DB {
	if _, ok := mapDatabaseClient[tenant]; !ok {

		//var tenantDatabaseClient *gorm.DB
		tenantDatabaseClient, err := database.NewTenantDatabaseClient(addr, tenant)
		if err == nil {
			mapDatabaseClient[tenant] = tenantDatabaseClient
		} else {
			log.Println("Can't create database client")
			return nil
		}

	}

	return mapDatabaseClient[tenant]
} */

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func GetState(client *gorm.DB) (bool, error) {
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
			var users []models.User
			result := client.Find(&users)
			if result.Error == nil {
				if result.RowsAffected >= 2 {
					state.Admin = true
					client.Save(&state)
				}
			}

		}
	}
	return err
}

func ChangeStateTenant(client *gorm.DB) (err error) {
	var state models.State
	err = client.First(&state).Error
	if err == nil {
		if !state.Admin {
			var admin models.Role
			err = client.Where("name = ?", "Administrator").First(&admin).Error
			if err == nil {
				var root models.Domain
				err = client.Where("name = ?", "Root").First(&root).Error
				if err == nil {
					var authorizations []models.Authorization
					result := client.Find(&authorizations)
					if result.Error == nil {
						if result.RowsAffected >= 2 {
							state.Admin = true
							client.Save(&state)
						}
					}
				}
			}
		}
	}
	return err
}

func GenerateHash() string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	concatenated := fmt.Sprint(random.Intn(100))
	sha512 := sha512.New()
	sha512.Write([]byte(concatenated))
	hash := base64.URLEncoding.EncodeToString(sha512.Sum(nil))
	return hash
}
