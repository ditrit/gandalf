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
	"gopkg.in/yaml.v2"

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
	evt.Timeout = 10000

	return
}

// IsExecAll : IsExecAll
func IsExecAll(mode os.FileMode) bool {
	return mode&0111 == 0111
}

// GetMaxVersion : GetMaxVersion
func GetMaxVersion(versions []int64) (maxversion int64) {
	maxversion = -1
	for _, version := range versions {
		if version > maxversion {
			maxversion = version
		}
	}
	return
}

// GetConnectorType : GetConnectorType
func GetConnectorType(connectorTypeName string, list []*models.ConnectorConfig) (result *models.ConnectorConfig) {
	for _, connectorType := range list {
		if connectorType.Name == connectorTypeName {
			result = connectorType
			break
		}
	}
	return result
}

// GetConnectorTypeConfigByVersion : GetConnectorTypeConfigByVersion
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

// GetConnectorCommand : GetConnectorCommand
func GetConnectorCommand(commandName string, list []models.ConnectorCommand) (result models.ConnectorCommand) {
	for _, command := range list {
		if command.Name == commandName {
			result = command
			break
		}
	}
	return result
}

// GetConnectorEvent : GetConnectorEvent
func GetConnectorEvent(eventName string, list []models.ConnectorEvent) (result models.ConnectorEvent) {
	for _, event := range list {
		if event.Name == eventName {
			result = event
			break
		}
	}
	return result
}

// ValidatePayload : Validate payload
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

// DownloadConfiguration : Download configuration from url
func DownloadConfiguration(url, ressource string) (connectorConfig *models.ConnectorConfig, err error) {

	resp, err := http.Get(url + ressource)
	if err != nil {
		log.Printf("err: %s", err)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(bodyBytes, &connectorConfig)
	if err != nil {
		log.Fatal(err)
	}

	return
}

// DownloadConfigurationsKeys : Download configurationsKeys from url
func DownloadConfigurationsKeys(url, ressource string) (body string, err error) {
	resp, err := http.Get(url + ressource)
	if err != nil {
		log.Printf("err: %s", err)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)

	return bodyString, nil
}

// DownloadWorkers : Download workers from url
func DownloadWorkers(url, filePath string) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("err: %s", err)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return
	}

	out, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		log.Printf("err: %s", err)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Printf("err: %s", err)
		return
	}

	return nil
}

// Unzip : Unzip file
func Unzip(zipPath string, dirPath string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		fpath := filepath.Join(dirPath, f.Name)

		if !strings.HasPrefix(fpath, filepath.Clean(dirPath)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

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

		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}
