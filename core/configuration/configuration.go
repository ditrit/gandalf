package configuration

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type configKey struct {
	value      *string // string key value (encoded as a string)
	component  string  // component Type
	shortName  string  // short parameter name (CLI)
	valType    string  // 'string' or 'integer' type of the value
	defaultVal string  // default value to use
	usage      string  // usage string (CLI)
	mandatory  bool    // is the value mandatory
}

var ConfigKeys = make(map[string]configKey)
var homePath = "/home/zippo/go/src"

//Set a string type key
func SetStringKeyConfig(componentType string, keyName string, shortName string, defaultValue string, usage string, mandatory bool) error {
	keyDef, exists := ConfigKeys[keyName]
	if exists {
		return errors.New("The key " + keyName + "is already defined ( for component " + keyDef.component + ")")
	}
	ConfigKeys[keyName] = configKey{new(string), componentType, shortName, "string", defaultValue, usage, mandatory}
	return nil
}

//Set an integer type key
func SetIntegerKeyConfig(componentType string, keyName string, shortName string, defaultValue int, usage string, mandatory bool) error {
	keyDef, exists := ConfigKeys[keyName]
	if exists {
		return errors.New("The key " + keyName + "is already defined ( for component " + keyDef.component + ")")
	}
	if defaultValue == -1 {
		ConfigKeys[keyName] = configKey{new(string), componentType, shortName, "integer", "", usage, mandatory}
	}
	ConfigKeys[keyName] = configKey{new(string), componentType, shortName, "integer", strconv.Itoa(defaultValue), usage, mandatory}
	return nil
}

//Get the value of a string type key
func GetStringConfig(keyName string) (string, error) {
	keyDef, exists := ConfigKeys[keyName]
	if !exists {
		return "", errors.New("config key " + keyName + "does not exist")
	}
	if keyDef.valType != "string" {
		return "", errors.New("The key " + keyName + " is of type string")
	}
	return *(keyDef.value), nil
}

//Get the value of an integer type key
func GetIntegerConfig(keyName string) (int, error) {
	keyDef, exists := ConfigKeys[keyName]
	if !exists {
		return -1, errors.New("config key " + keyName + "does not exist")
	}
	if keyDef.valType != "integer" {
		return -1, errors.New("The key " + keyName + " is of type integer")
	}
	intValue, err := strconv.Atoi(*(keyDef.value))
	if err != nil {
		return -1, errors.New("The value '" + *keyDef.value + "' provided for the key '" + keyName + "' can not be parsed as an integer.")
	}
	return intValue, nil
}

//launch all initiation functions
func InitMainConfigKeys() {
	InitCoreKeys()
	InitTenantKey()
	InitClusterKeys()
	InitAggregatorKeys()
	InitConnectorKeys()
}

//initiation of the core keys
func InitCoreKeys() {
	_ = SetStringKeyConfig("core", "config_dir", "f", homePath + "/gandalf/core/configuration/elements/", "path to the configuration directory", true)
	_ = SetStringKeyConfig("core", "logical_name", "l", "", "logical name of the component", true)
	_ = SetStringKeyConfig("core", "gandalf_type", "g", "", "launch mode (connector|aggregator|cluster)", true)
	_ = SetStringKeyConfig("core", "bind_address", "b", "", "Bind address", true)
	_ = SetStringKeyConfig("core", "certDir", "", homePath+"/gandalf/core/certs", "path of the TLS repository", false)
	_ = SetStringKeyConfig("core", "certPem", "", homePath+"/gandalf/core/certs/cert.pem", "path of the TLS certificate", false)
	_ = SetStringKeyConfig("core", "keyPem", "", homePath+"/gandalf/core/certs/key.pem", "path of the TLS private key", false)
	_ = SetStringKeyConfig("core", "caCertPem", "", homePath+"/gandalf/core/certs/ca.pem", "path of the CA certificate", false)
	_ = SetStringKeyConfig("core", "caKeyPem", "", homePath+"/gandalf/core/certs/cakey.pem", "path of the CA key", false)
	_ = SetStringKeyConfig("core", "gandalf_log", "", "/home/zippo/log/", "path of the log file", false)
}

//initiation of the tenant key for connectors and aggregators
func InitTenantKey() {
	_ = SetStringKeyConfig("connector/aggregator", "tenant", "t", "", "tenant of the component", true)
}

//initiation of the connector keys
func InitConnectorKeys() {
	_ = SetStringKeyConfig("connector", "connector_type", "y", "svn", "category of the connector", true)
	_ = SetStringKeyConfig("connector", "product_name", "p", "product1", "product of the connector", true)
	_ = SetStringKeyConfig("connector", "aggregators", "a", "address1:9800,address2:6400,address3", "aggregators addresses linked to the connector", true)
	_ = SetStringKeyConfig("connector", "gandalf_secret", "s", "/etc/gandalf/gandalfSecret", "path of the gandalf secret", true)
	_ = SetStringKeyConfig("connector", "product_url", "u", "url1,url2,url3", "product url list of the connector", false)
	_ = SetStringKeyConfig("connector", "workers", "w", "/etc/gandalf/workers", "path for the workers configuration", false)
	_ = SetStringKeyConfig("connector", "versions", "v", "1,2", "versions of a connector", true)
	_ = SetStringKeyConfig("connector", "grpc_bind_address", "r", "", "GRPC bind address", true)
	_ = SetIntegerKeyConfig("connector", "max_timeout", "m", 100, "maximum timeout of the connector", false)
}

//initiation of the aggregator keys
func InitAggregatorKeys() {
	_ = SetStringKeyConfig("aggregator", "clusters", "c", "address1:9800,address2:6300,address3", "clusters addresses linked to the aggregator", true)
}

//initiation of the cluster keys
func InitClusterKeys() {
	_ = SetStringKeyConfig("cluster", "cluster_join", "j", "", "cluster command (join)", false)
	_ = SetStringKeyConfig("cluster", "gandalf_db", "d", "/home/zippo/db", "path for the gandalf database", false)
}

//parse the configuration from the CLI parameters
func argParse(programName string, args []string) error {
	// parse CLI parameters
	flags := flag.NewFlagSet(programName, flag.ContinueOnError)
	var buf bytes.Buffer
	flags.SetOutput(&buf)
	for keyName := range ConfigKeys {
		keyDef := ConfigKeys[keyName]
		flags.StringVar(keyDef.value, keyName, "", keyDef.usage)
		if keyDef.shortName != "" {
			flags.StringVar(keyDef.value, keyDef.shortName, "", keyDef.usage)
		}
	}
	err := flags.Parse(args)
	if err != nil {
		return err
	}
	for keyName := range ConfigKeys {
		keyDef := ConfigKeys[keyName]
		if keyDef.valType == "integer" && *(keyDef.value) != "" {
			if _, err := strconv.Atoi(*(keyDef.value)); err != nil {
				return errors.New("error while parsing a CLI parameter: \n The CLI parameter for: " + keyName + " cannot be parsed as an integer using the value: " + *(keyDef.value))

			}
		}
	}
	return nil
}

//parse the configuration from the environment variables
func envParse() error {
	for keyName := range ConfigKeys {
		keyDef := ConfigKeys[keyName]
		strVal := os.Getenv("GANDALF_" + keyName)
		if keyDef.valType == "integer" && strVal != "" {
			if _, err := strconv.Atoi(strVal); err != nil {
				return errors.New("error while parsing an environment variable:\n The environment variable: GANDALF_" + keyName + " cannot be parsed as an integer using the value: " + strVal)
			}
		}
		if len(strVal) > 0 && *(keyDef.value) == "" {
			*(keyDef.value) = strVal
		}
	}
	return nil
}

//Set the map with all the values given in the yaml file
func yamlFileToMap() (map[interface{}]map[interface{}]string, error) {
	//Set a map from config yaml file
	keyDef := ConfigKeys["config_dir"]
	if *(keyDef.value) == "" {
		*(keyDef.value) = keyDef.defaultVal
	}
	file, err := os.Open(*(keyDef.value))
	if os.IsNotExist(err) {
		return nil, err
	}
	yamlMap := make(map[interface{}]map[interface{}]string)
	defer file.Close()

	//read all files in a given directory and unmarshal only yaml files
	list,_ := file.Readdirnames(0) // 0 to read all files and folders
	for _, fileName := range list {
		if filepath.Ext(*(keyDef.value) + fileName) == ".yaml" || filepath.Ext(*(keyDef.value) + fileName) == ".yml" {
			yamlFile, err := ioutil.ReadFile(*(keyDef.value) + fileName)
			err = yaml.Unmarshal(yamlFile, &yamlMap)
			if err != nil {
				return nil, errors.New("error while parsing the file")
			}
		}
	}

	return yamlMap, nil
}

//parse the configuration from the given yaml file
func yamlFileParse() error {
	tempMap, err := yamlFileToMap()
	if err != nil {
		return errors.New("error while parsing the file into a map")
	}

	//check if the all the keys in the yaml file are needed by the gandalf configuration
	for typeKey := range tempMap {
		for tempKeyName := range tempMap[typeKey] {
			keyName := fmt.Sprintf("%v", tempKeyName)
			_, found := ConfigKeys[keyName]
			if !found {
				return errors.New("The yaml key : " + keyName + " isn't needed by the gandalf configuration")
			}
		}
	}

	for keyName := range ConfigKeys {
		keyDef := ConfigKeys[keyName]
		if *(keyDef.value) == "" {
			*(keyDef.value) = tempMap[keyDef.component][keyName]
		}
		if keyDef.valType == "integer" && *(keyDef.value) != "" {
			if _, err := strconv.Atoi(*(keyDef.value)); err != nil {
				return errors.New("error while parsing a Yaml parameter:\n The Yaml parameter for: " + keyName + " cannot be parsed as an integer using the value: " + *(keyDef.value))
			}
		}
	}
	return nil
}

//parse the configuration with the default values
func defaultParse() error {
	for keyName := range ConfigKeys {
		keyDef := ConfigKeys[keyName]
		if *(keyDef.value) == "" {
			*(keyDef.value) = keyDef.defaultVal
		}
	}
	return nil
}

//Check if the configuration is valid (no invalid or missing values)
func IsConfigValid() error {
	gandalfType := *(ConfigKeys["gandalf_type"].value)
	if gandalfType != "connector" && gandalfType != "cluster" && gandalfType != "aggregator" {
		return errors.New("gandalf type isn't valid")
	}
	for keyName := range ConfigKeys {
		keyDef := ConfigKeys[keyName]
		if gandalfType == "connector" {
			if keyDef.component == "core" || keyDef.component == "connector" || keyDef.component == "connector/aggregator" {
				if keyDef.mandatory == true && *(keyDef.value) == "" {
					return errors.New(keyName + " is mandatory but has an invalid value")
				}
			}
		}
		if gandalfType == "aggregator" {
			if keyDef.component == "core" || keyDef.component == "aggregator" || keyDef.component == "connector/aggregator" {
				if keyDef.mandatory == true && *(keyDef.value) == "" {
					return errors.New(keyName + " is mandatory but has an invalid value")
				}
			}
		}
		if gandalfType == "cluster" {
			if keyDef.component == "core" || keyDef.component == "cluster" {
				if keyDef.mandatory == true && *(keyDef.value) == "" {
					return errors.New(keyName + " is mandatory but has an invalid value")
				}
			}
		}
	}
	return nil
}

//Set the list for all versions of a component
func GetVersionsList(strVal string) ([]int64, error) {
	var resultList []int64

	for _, val := range strings.Split(strVal, ",") {
		valint64, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return nil, err
		}
		resultList = append(resultList, valint64)
	}
	return resultList, nil
}

//set the map with the paths for the TLS keys
func setPathMap() map[string]string {
	certDir := *(ConfigKeys["certDir"].value)
	if filepath.IsAbs(certDir) {
		pathMap := map[string]string{
			"cert":  certDir + "/cert.pem",
			"key":   certDir + "/key.pem",
			"ca":    certDir + "/ca.pem",
			"cakey": certDir + "/cakey.pem",
		}
		return pathMap
	} else {
		pathMap := map[string]string{
			"cert":  *(ConfigKeys["certPem"].value),
			"key":   *(ConfigKeys["keyPem"].value),
			"ca":    *(ConfigKeys["caCertPem"].value),
			"cakey": *(ConfigKeys["caKeyPem"].value),
		}
		return pathMap
	}
}

//get the different TLS keys from the configuration
func GetTLS() (map[string][]byte, error) {
	pathMap := setPathMap()
	certMap := map[string][]byte{
		"cert":  []byte(""),
		"key":   []byte(""),
		"ca":    []byte(""),
		"cakey": []byte(""),
	}

	for certMapKey := range certMap {
		for pathKey := range pathMap {
			if certMapKey == pathKey {
				data, err := ioutil.ReadFile(pathMap[pathKey])
				if err != nil {
					certMap[certMapKey] = nil
					return certMap, err
				}
				certMap[certMapKey] = data
			}
		}
	}
	return certMap, nil
}

//Parse the configuration using the different parsing methods
func ParseConfig(programName string, args []string) error {
	err := argParse(programName, args)
	if err != nil {
		return err
	}
	err = envParse()
	if err != nil {
		return err
	}
	err = yamlFileParse()
	if err != nil {
		return err
	}
	err = defaultParse()
	return nil
}

//Launching configuration and testing if the configuration is valid
func ConfigMain(programName string, args []string) {
	InitMainConfigKeys()
	err := ParseConfig(programName, args)
	if err != nil {
		log.Fatalf("%v", err)
	}
	/*testTLS,err := GetTLS()
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(testTLS)*/
	err = IsConfigValid()
	if err != nil {
		log.Fatalf("%v", err)
	}
}
