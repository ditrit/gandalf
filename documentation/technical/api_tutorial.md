# Gandalf - Tutorial creation of new route for the API
## Structure

The api is divided into several files

gandalf/core/aggregator/api/api_[models]. go Contains functions (create/update/list/read/delete/etc...) used by router routes.go

gandalf/core/aggregator/api/middleware.go Contains the IsAuthorized() function that allows verification of the JWT token. Currently only user authentication is checked.

> TODO: Management of authorisations in the IsAuthorized() function

gandalf/core/aggregator/api/logger.go Contains code for logging API calls

gandalf/core/aggregator/api/routers.go Contains the routes and routers used by the API

gandalf/core/aggregator/api/apiServer.go Contains the code for the API server

gandalf/core/aggregator/api/dao/[models]DAO.go Contains the ORM GORM binding functions.

gandalf/core/models Contains gandalf data structures Note: Most data structures Gandalf uses the Model.go structure to add fields (ID, CreatedAt, UpdatedAt, DeletedAt) to the structure.

## Creation of a new API route

+ Step 1: Create the models if necessary (gandalf/core/models) and update the file (gandalf/core/cluster/database/databaseConnection.go)
+ Step 2: Creation of the DAO if necessary (gandalf/core/aggregator/api/dao)
+ Step 3: Create the api_[models] file. go (gandalf/core/aggregator/api)
+ Step 4: Write the functions required for operation in the api_[models] file. go
+ Step 5: Add routes to the gandalf/core/aggregator/routers.go file using the functions previously written in the api_[models] file. go
+ Step 6: Update Swagger documentation (https://app.swaggerhub.com/apis/ditrit/Gandalf/1.0.0-oas3) and the one located in gandalf/core/aggregator/api/swagger.yaml

## Example of a new Toto API route:

+ Step 1: Creation of toto.go models in gandalf/core/models and AutoMigrate update of "InitGandalfDatabase" and "InitTenantDatabase" functions in gandalf/core/cluster/database/databaseConnection.go

```
package models

type Toto struct {
	Model
	Name             string `gorm:"unique;not null"`
}
```

```
// InitGandalfDatabase : Gandalf database init.
func (dc DatabaseConnection) InitGandalfDatabase(gandalfDatabaseClient *gorm.DB, logicalName, bindAddress string) (login []string, password []string, err error) {
	gandalfDatabaseClient.AutoMigrate(&models.State{}, &models.Event{}, &models.Tenant{}, &models.SecretAssignement{},
		&models.Command{}, &models.Authorization{}, &models.Role{}, &models.User{}, &models.Domain{},
		&models.Pivot{}, &models.ProductConnector{}, &models.ConnectorProduct{}, &models.Key{}, &models.CommandType{}, &models.EventType{},
		&models.ResourceType{}, &models.Resource{}, &models.KeyValue{}, &models.LogicalComponent{}, &models.EventTypeToPoll{}, &models.Heartbeat{},
		&models.Tag{}, &models.Library{}, &models.Product{}, &models.Environment{}, &models.EnvironmentType{}, &models.Toto{})
...
```

+ Step 2 : Creation of DAO totoDAO.go dans gandalf/core/aggregator/api/dao

```
package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/google/uuid"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListToto(database *gorm.DB) (totos []models.Toto, err error) {
	err = database.Find(&totos).Error

	return
}

func CreateToto(database *gorm.DB, toto *models.Toto) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Create(&toto).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadToto(database *gorm.DB, id uuid.UUID) (toto models.Toto, err error) {
	err = database.Where("id = ?", id).First(&toto).Error

	return
}

func UpdateToto(database *gorm.DB, toto models.Toto) (err error) {
	err = database.Save(&toto).Error

	return
}

func DeleteToto(database *gorm.DB, id uuid.UUID) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var toto models.Toto
			err = database.Where("id = ?", id).First(&toto).Error
			if err == nil {
				err = database.Unscoped().Delete(&toto).Error
			}
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
```
+ Step 3 and 4: Creating the file api_toto.go in gandalf/core/aggregator/api

```
package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/ditrit/gandalf/core/aggregator/api/dao"
	"github.com/ditrit/gandalf/core/models"
	"github.com/google/uuid"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/gorilla/mux"
)

func CreateToto(w http.ResponseWriter, r *http.Request) {

	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		var toto *models.Toto
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&toto); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		if err := dao.CreateToto(database, toto); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, toto)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func DeleteToto(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := uuid.Parse(vars["totoId"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}

		if err := dao.DeleteToto(database, id); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func GetTotoById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := uuid.Parse(vars["totoId"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
			return
		}

		var toto models.Toto
		if toto, err = dao.ReadToto(database, id); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "User not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, toto)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func ListToto(w http.ResponseWriter, r *http.Request) {

	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		totos, err := dao.ListToto(database)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, totos)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func UpdateToto(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := uuid.Parse(vars["totoId"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
			return
		}

		var toto models.Toto
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&toto); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
			return
		}
		defer r.Body.Close()
		toto.ID = id

		if err := dao.UpdateToto(database, toto); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, toto)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}
```
+ Step 5: Adding routes to the routers.go file located in gandalf/core/aggregator/api
```
	Route{
		"CreateToto",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/toto",
		IsAuthorized(CreateToto),
	},

	Route{
		"DeleteToto",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/toto/{totoId:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}}",
		IsAuthorized(DeleteToto),
	},

	Route{
		"GetTotoById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/toto/{totoId:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}}",
		IsAuthorized(GetTotoById),
	},

	Route{
		"ListToto",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/toto",
		IsAuthorized(ListToto),
	},

	Route{
		"UpdateToto",
		strings.ToUpper("Put"),
		"/ditrit/Gandalf/1.0.0/toto/{totoId}",
		IsAuthorized(UpdateToto),
	},
```
+ Step 6: Swagger documentation upgrade

+ Step 7: Enjoy
