//Package utils :
package utils

import (
	"archive/zip"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

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

// GetMaxVersion : GetMaxVersion
func GetMaxVersion(versions []models.Version) (maxversion models.Version) {
	maxversion = models.Version{Major: 0, Minor: 0}
	for _, version := range versions {
		if version.Major > maxversion.Major {
			maxversion = version
		} else if version.Major == maxversion.Major {
			if version.Minor > maxversion.Minor {
				maxversion = version
			}
		}
	}
	return
}

/* // GetConnectorType : GetConnectorType
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
func GetConnectorTypeConfigByVersion(major int8, list []*models.ConnectorConfig) (result *models.ConnectorConfig) {
	for _, connectorConfig := range list {
		if connectorConfig.Major == major {
			result = connectorConfig
			break
		}
	}

	return result
}

// GetConnectorCommand : GetConnectorCommand
func GetConnectorCommand(commandName string, list []models.Object) (result models.Object) {
	for _, command := range list {
		if command.Name == commandName {
			result = command
			break
		}
	}
	return result
}

// GetConnectorEvent : GetConnectorEvent
func GetConnectorEvent(eventName string, list []models.Object) (result models.Object) {
	for _, event := range list {
		if event.Name == eventName {
			result = event
			break
		}
	}
	return result
} */

// GetConnectorCommand : GetConnectorCommand
func GetConnectorCommandType(commandName string, list []models.CommandType) (result models.CommandType) {
	for _, command := range list {
		if command.Name == commandName {
			result = command
			break
		}
	}
	return result
}

// GetConnectorEvent : GetConnectorEvent
func GetConnectorEventType(eventName string, list []models.EventType) (result models.EventType) {
	for _, event := range list {
		if event.Name == eventName {
			result = event
			break
		}
	}
	return result
}

// GetPivotByVersion : GetPivotByVersion
func GetPivotByVersion(major, minor int8, pivots []*models.Pivot) (result *models.Pivot) {
	for _, pivot := range pivots {
		if pivot.Major == major && pivot.Minor == minor {
			result = pivot
			break
		}
	}

	return
}

// GetPivotByVersion : GetPivotByVersion
func GetConnectorProductByVersion(major, minor int8, productConnectors []*models.ProductConnector) (result *models.ProductConnector) {
	for _, productConnector := range productConnectors {
		if productConnector.Major == major && productConnector.Minor == minor {
			result = productConnector
			break
		}
	}

	return
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

// DownloadPivot : Download pivot from url
func DownloadPivot(url, ressource string) (pivot *models.Pivot, err error) {

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

	err = yaml.Unmarshal(bodyBytes, &pivot)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	return
}

// DownloadConnectorProduct : Download connector product from url
func DownloadProductConnector(url, ressource string) (productConnector *models.ProductConnector, err error) {

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

	err = yaml.Unmarshal(bodyBytes, &productConnector)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	return
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
		fmt.Println(err)
		log.Fatal(err)
	}

	return
}

// DownloadConfigurationsKeys : Download configurationsKeys from url
func DownloadConfigurationsKeys(url, ressource string) (body string, err error) {

	resp, err := http.Get(url + ressource)
	if err != nil {
		fmt.Println(err)
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

// DownloadVersions : Download versions from url
func DownloadVersions(url, ressource string) (versions []string, err error) {
	fmt.Println("url")
	fmt.Println(url)
	fmt.Println(ressource)
	fmt.Println(url + ressource)
	fmt.Println("url1")
	resp, err := http.Get(url + ressource)
	if err != nil {
		log.Printf("err: %s", err)
		fmt.Printf("err: %s", err)
		return
	}
	fmt.Println("url2")
	fmt.Println(resp)
	fmt.Println(resp.Body)
	fmt.Println(resp.StatusCode)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("url200")
		return
	}
	fmt.Println("url3")

	fmt.Println("resp.Body")
	fmt.Println(resp.Body)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("string(bodyBytes)")
	fmt.Println(string(bodyBytes))
	err = yaml.Unmarshal(bodyBytes, &versions)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	return
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

// IsExecAll : IsExecAll
func IsExecAll(mode os.FileMode) bool {
	return mode&0111 == 0111
}

// CheckFileExist: CheckFileExist
func CheckFileExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func CheckFileExistAndIsExecAll(path string) bool {
	if fileState, err := os.Stat(path); err == nil {
		if IsExecAll(fileState.Mode()) {
			return true
		}
		return false
	}
	return false
}

func GenerateHash(logicalName string) string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	concatenated := fmt.Sprint(logicalName, random.Intn(100))
	sha512 := sha512.New()
	sha512.Write([]byte(concatenated))
	hash := base64.URLEncoding.EncodeToString(sha512.Sum(nil))
	return hash
}
