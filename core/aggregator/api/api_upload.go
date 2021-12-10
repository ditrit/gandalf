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
	fileDir := "/var/lib/gandalf/files/"
	filePath := fileDir + fileId
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		// TODO: Do not must crash + logs
		fmt.Println(err)
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
		}
	}

	dst, err := os.Create(fileDir + fileId)
	defer dst.Close()
	if err != nil {
		// TODO: Do not must crash + logs
		fmt.Println(err)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		fmt.Println(err)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}
