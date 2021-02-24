package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ditrit/gandalf/core/cluster/database"

	"github.com/ditrit/gandalf/core/cluster/api/dao"
	"github.com/ditrit/gandalf/core/cluster/api/utils"
	"github.com/ditrit/gandalf/core/models"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

// ConfigurationAggregatorController :
type ConfigurationAggregatorController struct {
	databaseConnection *database.DatabaseConnection
}

// NewConfigurationController :
func NewConfigurationAggregatorController(databaseConnection *database.DatabaseConnection) (configurationAggregatorController *ConfigurationAggregatorController) {
	configurationAggregatorController = new(ConfigurationAggregatorController)
	configurationAggregatorController.databaseConnection = databaseConnection

	return
}

func (cc ConfigurationAggregatorController) Upload(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := cc.databaseConnection.GetDatabaseClientByTenant(tenant)
	if database != nil {
		fmt.Println("File Upload Endpoint Hit")

		r.ParseMultipartForm(10 << 20)

		file, handler, err := r.FormFile("myFile")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		var configurationAggregator *models.ConfigurationLogicalAggregator
		err = yaml.Unmarshal(fileBytes, &configurationAggregator)
		if err != nil {
			fmt.Println(err)
		}

		database.Save(&configurationAggregator)

		fmt.Fprintf(w, "Successfully Uploaded File\n")
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// List :
func (cc ConfigurationAggregatorController) List(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := cc.databaseConnection.GetDatabaseClientByTenant(tenant)
	if database != nil {
		configurationAggregators, err := dao.ListConfigurationAggregator(database)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, configurationAggregators)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}

}

// Create :
func (cc ConfigurationAggregatorController) Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := cc.databaseConnection.GetDatabaseClientByTenant(tenant)
	if database != nil {

		var configurationAggregator models.ConfigurationLogicalAggregator
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&configurationAggregator); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		if err := dao.CreateConfigurationAggregator(database, configurationAggregator); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, configurationAggregator)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Read :
func (cc ConfigurationAggregatorController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := cc.databaseConnection.GetDatabaseClientByTenant(tenant)
	if database != nil {

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var configurationAggregator models.ConfigurationLogicalAggregator
		if configurationAggregator, err = dao.ReadConfigurationAggregator(database, id); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, configurationAggregator)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Update :
func (cc ConfigurationAggregatorController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := cc.databaseConnection.GetDatabaseClientByTenant(tenant)
	if database != nil {

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var configurationAggregator models.ConfigurationLogicalAggregator
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&configurationAggregator); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
			return
		}
		defer r.Body.Close()
		configurationAggregator.ID = uint(id)

		if err := dao.UpdateConfigurationAggregator(database, configurationAggregator); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, configurationAggregator)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Delete :
func (cc ConfigurationAggregatorController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := cc.databaseConnection.GetDatabaseClientByTenant(tenant)
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}

		if err = dao.DeleteConfigurationAggregator(database, id); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}
