package utils

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gandalf/core/cluster/utils"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ditrit/gandalf/core/models"
	"gopkg.in/yaml.v2"

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

func GetPivot(client *gorm.DB, baseurl, componentType string, version models.Version) (*models.Pivot, error) {
	var pivot *models.Pivot
	err := client.Where("name = ? and major = ? and minor = ?", componentType, version.Major, version.Minor).Preload("ResourceTypes").Preload("CommandTypes").Preload("EventTypes").Preload("Keys").First(&pivot).Error
	if err != nil {
		pivot, _ = utils.DownloadPivot(baseurl, "/configurations/"+strings.ToLower(componentType)+"/"+strconv.Itoa(int(version.Major))+"_"+strconv.Itoa(int(version.Minor))+"_pivot.yaml")
	}

	client.Create(&pivot)

	return pivot, nil
}

// DownloadPivot : Download pivot from url
func DownloadPivot(url, ressource string) (pivot *models.Pivot, err error) {

	resp, err := http.Get(url + ressource)
	if err != nil {
		log.Printf("err: %s", err)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(bodyBytes, &pivot)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	return
}

func SaveLogicalComponent(client *gorm.DB, logicalName, repositoryURL string, pivot *models.Pivot) (*models.LogicalComponent, error) {
	logicalComponent := new(models.LogicalComponent)
	logicalComponent.LogicalName = logicalName
	logicalComponent.Type = "aggregator"
	logicalComponent.Pivot = pivot
	var keyValues []models.KeyValue
	for _, key := range pivot.Keys {
		keyValue := new(models.KeyValue)
		switch key.Name {
		case "repository_url":
			keyValue.Value = repositoryURL
			keyValue.Key = key
			keyValues = append(keyValues, *keyValue)
		}

	}

	logicalComponent.KeyValues = keyValues

	client.Create(&logicalComponent)

	return logicalComponent, nil
}
