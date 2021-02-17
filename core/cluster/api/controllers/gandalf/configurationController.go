package gandalf

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ditrit/gandalf/core/cluster/database"

	"github.com/ditrit/gandalf/core/cluster/api/utils"
	"github.com/ditrit/gandalf/core/models"

	"github.com/ditrit/gandalf/core/cluster/api/dao"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

// ConfigurationController :
type ConfigurationController struct {
	databaseConnection *database.DatabaseConnection
}

// NewConfigurationController :
func NewConfigurationController(databaseConnection *database.DatabaseConnection) (configurationController *ConfigurationController) {
	configurationController = new(ConfigurationController)
	configurationController.databaseConnection = databaseConnection

	return
}

func (cc ConfigurationController) Upload(w http.ResponseWriter, r *http.Request) {
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

	var configurationConfigurationCluster *models.ConfigurationLogicalCluster
	err = yaml.Unmarshal(fileBytes, &configurationConfigurationCluster)
	if err != nil {
		fmt.Println(err)
	}

	cc.databaseConnection.GetGandalfDatabaseClient().Save(&configurationConfigurationCluster)

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

// List :
func (cc ConfigurationController) List(w http.ResponseWriter, r *http.Request) {
	configurationCluster, err := dao.ListConfigurationCluster(cc.databaseConnection.GetGandalfDatabaseClient())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, configurationCluster)
}

// Create :
func (cc ConfigurationController) Create(w http.ResponseWriter, r *http.Request) {
	var configurationCluster models.ConfigurationLogicalCluster
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&configurationCluster); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := dao.CreateConfigurationCluster(cc.databaseConnection.GetGandalfDatabaseClient(), configurationCluster); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, configurationCluster)
}

// Read :
func (cc ConfigurationController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var configurationCluster models.ConfigurationLogicalCluster
	if configurationCluster, err = dao.ReadConfigurationCluster(cc.databaseConnection.GetGandalfDatabaseClient(), id); err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.RespondWithError(w, http.StatusNotFound, "Product not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, configurationCluster)
}

// Update :
func (cc ConfigurationController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var configurationCluster models.ConfigurationLogicalCluster
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&configurationCluster); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	configurationCluster.ID = uint(id)

	if err := dao.UpdateConfigurationCluster(cc.databaseConnection.GetGandalfDatabaseClient(), configurationCluster); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, configurationCluster)
}

// Delete :
func (cc ConfigurationController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	if err := dao.DeleteConfigurationCluster(cc.databaseConnection.GetGandalfDatabaseClient(), id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
