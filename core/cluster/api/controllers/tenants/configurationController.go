package tenants

import (
	"database/sql"
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

// ConfigurationController :
type ConfigurationController struct {
	mapDatabase map[string]*gorm.DB
}

// NewConfigurationController :
func NewConfigurationController(mapDatabase map[string]*gorm.DB) (configurationController *ConfigurationController) {
	configurationController = new(ConfigurationController)
	configurationController.mapDatabase = mapDatabase

	return
}

func (cc ConfigurationController) UploadConnector(w http.ResponseWriter, r *http.Request) {
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

		var configurationConnector *models.ConfigurationConnector
		err = yaml.Unmarshal(fileBytes, &configurationConnector)
		if err != nil {
			fmt.Println(err)
		}

		database.Save(&configurationConnector)

		fmt.Fprintf(w, "Successfully Uploaded File\n")
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func (cc ConfigurationController) UploadAggregator(w http.ResponseWriter, r *http.Request) {
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

		var configurationAggregator *models.ConfigurationAggregator
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

// Read :
func (cc ConfigurationController) ReadConnector(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(cc.mapDatabase, tenant)
	if database != nil {

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var cluster models.Cluster
		if cluster, err = dao.ReadCluster(database, id); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, cluster)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Read :
func (cc ConfigurationController) ReadAggregator(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(cc.mapDatabase, tenant)
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var cluster models.Cluster
		if cluster, err = dao.ReadCluster(database, id); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, cluster)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}
