package tenants

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ditrit/gandalf/core/cluster/api/dao"
	"github.com/ditrit/gandalf/core/cluster/api/utils"
	"github.com/ditrit/gandalf/core/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
)

// ConfigurationConnectorController :
type ConfigurationConnectorController struct {
	mapDatabase map[string]*gorm.DB
}

// NewConfigurationConnectorController :
func NewConfigurationConnectorController(mapDatabase map[string]*gorm.DB) (configurationConnectorController *ConfigurationConnectorController) {
	configurationConnectorController = new(ConfigurationConnectorController)
	configurationConnectorController.mapDatabase = mapDatabase

	return
}

func (cc ConfigurationConnectorController) Upload(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(cc.mapDatabase, tenant)
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

		var configurationConfigurationConnector *models.ConfigurationLogicalConnector
		err = yaml.Unmarshal(fileBytes, &configurationConfigurationConnector)
		if err != nil {
			fmt.Println(err)
		}

		database.Save(&configurationConfigurationConnector)

		fmt.Fprintf(w, "Successfully Uploaded File\n")
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// List :
func (cc ConfigurationConnectorController) List(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(cc.mapDatabase, tenant)
	if database != nil {
		configurationConnectors, err := dao.ListConfigurationConnector(database)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, configurationConnectors)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Create :
func (cc ConfigurationConnectorController) Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(cc.mapDatabase, tenant)
	if database != nil {
		var configurationConnector models.ConfigurationLogicalConnector
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&configurationConnector); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		if err := dao.CreateConfigurationConnector(database, configurationConnector); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, configurationConnector)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Read :
func (cc ConfigurationConnectorController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(cc.mapDatabase, tenant)
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var configurationConnector models.ConfigurationLogicalConnector
		if configurationConnector, err = dao.ReadConfigurationConnector(database, id); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, configurationConnector)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Update :
func (cc ConfigurationConnectorController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(cc.mapDatabase, tenant)
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var configurationConnector models.ConfigurationLogicalConnector
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&configurationConnector); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
			return
		}
		defer r.Body.Close()
		configurationConnector.ID = uint(id)

		if err := dao.UpdateConfigurationConnector(database, configurationConnector); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, configurationConnector)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Delete :
func (cc ConfigurationConnectorController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(cc.mapDatabase, tenant)
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}

		if err := dao.DeleteConfigurationConnector(database, id); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}
