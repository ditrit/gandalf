package api

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	cmodels "github.com/ditrit/gandalf/core/configuration/models"

	"github.com/gorilla/mux"
)

func GetFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileId := vars["fileId"]

	// TODO: must be configurable, check if directory already exists
	configuration := utils.Shoset.Context["configuration"].(*cmodels.ConfigurationAggregator)
	fileDir := configuration.GetAPIPath()
	filePath := fileDir + fileId
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		// TODO: Do not must crash + logs
		fmt.Println(err)
		utils.RespondWithError(w, http.StatusBadRequest, "File doesn't exist")
		return
	}

	http.ServeFile(w, r, filePath)
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)
	vars := mux.Vars(r)
	fileId := vars["fileId"]

	// Get handler for filename, size and headers
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		utils.RespondWithError(w, http.StatusBadRequest, "Error retrieving the file")
		return
	}
	defer file.Close()

	configuration := utils.Shoset.Context["configuration"].(*cmodels.ConfigurationAggregator)
	fileDir := configuration.GetAPIPath()
	_, err = os.Stat(fileDir)
	if os.IsNotExist(err) {
		err = os.Mkdir(fileDir, 0666)
		if err != nil {
			// TODO: Do not must crash + logs
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusInternalServerError, "Error retrieving the directory")
		}
	}

	dst, err := os.Create(fileDir + fileId)
	defer dst.Close()
	if err != nil {
		// TODO: Do not must crash + logs
		fmt.Println(err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Error creating the new file")
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		fmt.Println(err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Error copying the file")

		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "Successfully Uploaded File"})
}
