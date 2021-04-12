package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ditrit/gandalf/core/aggregator/database"

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
	database := lc.databaseConnection.GetTenantDatabaseClient()
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

		var logicalComponent *models.LogicalComponent
		err = yaml.Unmarshal(fileBytes, &logicalComponent)
		if err != nil {
			fmt.Println(err)
		}

		//TODO ADD VALIDATION

		database.Save(&logicalComponent)

		fmt.Fprintf(w, "Successfully Uploaded File\n")
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}
