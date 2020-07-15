//Package utils :
package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/ditrit/gandalf/core/models"

	"github.com/ditrit/shoset/msg"

	"github.com/xeipuuv/gojsonschema"
)

// CreateValidationEvent : Connector create validation event functions.
func CreateValidationEvent(command msg.Command, tenant string) (evt *msg.Event) {
	var tab = map[string]string{
		"topic":         command.GetCommand(),
		"event":         "ON_GOING",
		"payload":       "",
		"referenceUUID": command.GetUUID()}

	evt = msg.NewEvent(tab)
	evt.Tenant = tenant
	evt.Timeout = 100000

	return
}

//IsExecAll
func IsExecAll(mode os.FileMode) bool {
	return mode&0111 == 0111
}

//GetMaxVersion
func GetMaxVersion(versions []int64) (maxversion int64) {
	maxversion = -1
	for _, version := range versions {
		if version > maxversion {
			maxversion = version
		}
	}
	return
}

//GetConnectorType
func GetConnectorType(connectorTypeName string, list []*models.ConnectorConfig) (result *models.ConnectorConfig) {
	for _, connectorType := range list {
		if connectorType.Name == connectorTypeName {
			result = connectorType
			break
		}
	}
	return result
}

//GetConnectorTypeConfigByVersion
func GetConnectorTypeConfigByVersion(version int64, list []*models.ConnectorConfig) (result *models.ConnectorConfig) {
	if version == 0 {
		result = nil
	} else {
		for _, connectorConfig := range list {
			if int64(connectorConfig.Version) == version {
				result = connectorConfig
				break
			}
		}
	}

	return result
}

//TODO REVOIR INTERFACE
//GetConnectorTypeCommand
func GetConnectorTypeCommand(commandName string, list []models.ConnectorTypeCommand) (result models.ConnectorTypeCommand) {
	for _, command := range list {
		if command.Name == commandName {
			result = command
			break
		}
	}
	return result
}

//TODO REVOIR INTERFACE
//GetConnectorTypeEvent
func GetConnectorTypeEvent(eventName string, list []models.ConnectorTypeEvent) (result models.ConnectorTypeEvent) {
	for _, event := range list {
		if event.Name == eventName {
			result = event
			break
		}
	}
	return result
}

//ValidatePayload
func ValidatePayload(payload, payloadSchema string) (result bool) {
	result = false

	payloadloader := gojsonschema.NewStringLoader(payload)
	payloadSchemaloader := gojsonschema.NewStringLoader(payloadSchema)

	validate, err := gojsonschema.Validate(payloadSchemaloader, payloadloader)
	if err != nil {
		log.Printf("Error on validation payload : %s", err)
	} else {
		if validate.Valid() {
			result = true
		}
	}
	return result

}

//TODO REVOIR
//DownloadWorkers
func DownloadConfigurationsKeys(url, ressource string) (body string, err error) {
	// Create the file
	resp, err := http.Get(url + ressource)
	if err != nil {
		log.Printf("err: %s", err)
		return
	}

	defer resp.Body.Close()
	fmt.Println("status", resp.Status)
	if resp.StatusCode != 200 {
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println("bodystring")
	fmt.Println(bodyString)
	return bodyString, nil
}

//DownloadWorkers
func DownloadWorkers(url, filePath string) (err error) {
	// Create the file
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("err: %s", err)
		return
	}

	defer resp.Body.Close()
	fmt.Println("status", resp.Status)
	if resp.StatusCode != 200 {
		return
	}

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		log.Printf("err: %s", err)
		return
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Printf("err: %s", err)
		return
	}

	return nil
}

//Unzip
func Unzip(zipPath string, dirPath string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dirPath, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dirPath)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}
