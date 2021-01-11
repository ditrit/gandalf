package configuration

import (
	"database/sql"
	"fmt"
	"gandalf/core/cluster/api/dao"
	"gandalf/core/cluster/api/utils"
	"gandalf/core/models"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
)

// ConfigurationController :
type ConfigurationController struct {
	gandalfDatabase *gorm.DB
}

// NewConfigurationController :
func NewConfigurationController(gandalfDatabase *gorm.DB) (configurationController *ConfigurationController) {
	configurationController = new(ConfigurationController)
	configurationController.gandalfDatabase = gandalfDatabase

	return
}

func UploadClusterConfiguration(w http.ResponseWriter, r *http.Request) {
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

	var configurationCluster *models.ConfigurationCluster
	err = yaml.Unmarshal(fileBytes, &configurationCluster)
	if err != nil {
		fmt.Println(err)
	}

	gandalfDatabase.Save(&configurationCluster)

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

// Read :
func (cc ConfigurationController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var cluster models.Cluster
	if cluster, err = dao.ReadCluster(cc.gandalfDatabase, id); err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.RespondWithError(w, http.StatusNotFound, "Product not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, cluster)
}
