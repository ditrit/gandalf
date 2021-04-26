package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ditrit/gandalf/core/aggregator/database"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/ditrit/gandalf/core/models"

	"gopkg.in/yaml.v2"
)

// ConfigurationAggregatorController :
type LogicalComponentController struct {
	databaseConnection *database.DatabaseConnection
}

// NewConfigurationController :
func NewLogicalComponentController(databaseConnection *database.DatabaseConnection) (logicalComponentController *LogicalComponentController) {
	logicalComponentController = new(LogicalComponentController)
	logicalComponentController.databaseConnection = databaseConnection

	return
}

func (lc LogicalComponentController) Upload(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	typeComponent := vars["type"]
	tenant := vars["tenant"]
	major, err := strconv.Atoi(vars["major"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid major ")
		return
	}
	minor, err := strconv.Atoi(vars["minor"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid minor ")
		return
	}
	version := models.Version{Major: int8(major), Minor: int8(minor)}
	database := lc.databaseConnection.GetDatabaseClientByTenant(tenant)
	if database != nil {

		pivot, err := utils.GetPivot(database, lc.databaseConnection.GetLogicalComponent().GetKeyValueByKey("repository_url").Value, typeComponent, version)
		if err == nil {
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

			var logicalComponent *models.LogicalComponent
			err = yaml.Unmarshal(fileBytes, &logicalComponent)
			if err != nil {
				fmt.Println(err)
			}

			//TODO ADD VALIDATION
			SaveLogicalComponent(database, logicalComponent, pivot)

			fmt.Fprintf(w, "Successfully Uploaded File\n")
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, "pivot not found")
			return
		}
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func SaveLogicalComponent(client *gorm.DB, logicalComponent *models.LogicalComponent, pivot models.Pivot) error {
	fmt.Println("test")
	fmt.Println(logicalComponent.KeyValues)
	logicalComponent.Pivot = pivot
	for _, keyValue := range logicalComponent.KeyValues {
		for _, key := range pivot.Keys {
			if keyValue.Key.Name == key.Name {
				keyValue.Key = key
				break
			}
		}
	}
	fmt.Println("logicalComponent")
	fmt.Println(logicalComponent)
	err := client.Create(&logicalComponent).Error
	fmt.Println(err)
	return nil
}
