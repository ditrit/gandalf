package configuration

import (
	"errors"
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type configKey struct {
	value      *string // string key value (encoded as a string)
	component  string  // component Type
	shortName  string  // short parameter name (CLI)
	valType    string  // 'string' or 'integer' type of the value
	defaultVal string  // default value to use
	usage      string  // usage string (CLI)
}

var ConfigKeys = make(map[string]configKey)

// SetStringKeyConfig :
func SetStringKeyConfig(componentType string, keyName string, shortName string, defaultValue string, usage string) error {
	keyDef, exists := ConfigKeys[keyName]
	if exists {
		return errors.New("The key " + keyName + "is already defined ( for component " + keyDef.component + ")")
	}
	ConfigKeys[keyName] = configKey{new(string), componentType, shortName, "string", defaultValue, usage}
	return nil
}

// SetIntegerKeyConfig :
func SetIntegerKeyConfig(componentType string, keyName string, shortName string, defaultValue int, usage string) error {
	keyDef, exists := ConfigKeys[keyName]
	if exists {
		return errors.New("The key " + keyName + "is already defined ( for component " + keyDef.component + ")")
	}
	ConfigKeys[keyName] = configKey{new(string), componentType, shortName, "integer", strconv.Itoa(defaultValue), usage}
	return nil
}

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

func InitMainConfigKeys() {
	InitCoreKeys()
	InitConnectorKeys()
	InitAggregatorKeys()
	InitClusterKeys()
}

func InitCoreKeys() {
	_ = SetStringKeyConfig("core", "config_file", "f", "configuration/elements/gandalf.yaml", "path to the configuration file")
	_ = SetStringKeyConfig("core", "logical_name", "l", "logical name", "logical name of the component")
	_ = SetStringKeyConfig("core", "gandalf_type", "g", "gandalf type", "launch mode (connector|aggregator|cluster)")
	_ = SetStringKeyConfig("core", "cert_pem", "", "/etc/gandalf/cert/cert.pem", "path of the TLS certificate")
	_ = SetStringKeyConfig("core", "key_pem", "", "/etc/gandalf/cert/key.pem", "path of the TLS private key")
}

func InitConnectorKeys(){
	_ = SetStringKeyConfig("connector","tenant","t","tenant1","tenant of the connector")
	_ = SetStringKeyConfig("connector","category","c","svn","category of the connector")
	_ = SetStringKeyConfig("connector", "product","p","product1","product of the connector")
	_ = SetStringKeyConfig("connector","aggregators", "a","address1:9800,address2:6400,address3","aggregators addresses linked to the connector")
	_ = SetStringKeyConfig("connector","gandalf_secret","s","/etc/gandalf/gandalfSecret","path of the gandalf secret")
	_ = SetStringKeyConfig("connector","product_url","u","url1,url2,url3","product url list of the connector")
	_ = SetStringKeyConfig("connector","connector_log","","/etc/gandalf/log","path of the log file")
	_ = SetIntegerKeyConfig("connector","max_timeout","",100,"maximum timeout of the connector")
}

func InitAggregatorKeys(){
	_ = SetStringKeyConfig("aggregator","aggregator_tenant","","tenant1","tenant of the aggregator")
	_ = SetStringKeyConfig("aggregator","cluster","","address1[:9800],address2[:6300],address3","clusters addresses linked to the aggregator")
	_ = SetStringKeyConfig("aggregator","aggregator_log","","/etc/gandalf/log","path of the log file")
}

func InitClusterKeys(){
	_ = SetStringKeyConfig("cluster","join","j","clusterAddress","link the cluster member to another one")
	_ = SetStringKeyConfig("cluster","cluster_log","","/etc/gandalf/log","path of the log file")
	_ = SetStringKeyConfig("cluster","gandalf_db","d","pathToTheDB","path for the gandalf database")
}

func argParse() {
	// parse CLI parameters
	for keyName := range ConfigKeys {
		keyDef := ConfigKeys[keyName]
		flag.StringVar(keyDef.value, keyName, "", keyDef.usage)
		if keyDef.shortName != "" {
			flag.StringVar(keyDef.value, keyDef.shortName, "", keyDef.usage)
		}
	}
	flag.Parse()
	for keyName := range ConfigKeys {
		keyDef := ConfigKeys[keyName]
		if keyDef.valType == "integer" && *(keyDef.value) != "" {
			if _, err := strconv.Atoi(*(keyDef.value)); err != nil {
				newErr := errors.New("The CLI parameter for: " + keyName + " cannot be parsed as an integer using the value: " + *(keyDef.value))
				log.Fatalf("error while parsing a CLI parameter: %v",newErr)
			}
		}
	}
}

func tempEnvVarSet() {
	//temporary environment variables setter
	_ = os.Setenv("GANDALF_connectorFlag", "testENV")
	_ = os.Setenv("GANDALF_connector2", "12")
}

func envParse() {
	// parse environment variables
	for keyName := range ConfigKeys {
		keyDef := ConfigKeys[keyName]
		strVal := os.Getenv("GANDALF_" + keyName)
		if keyDef.valType == "integer" && strVal != "" {
			if _, err := strconv.Atoi(strVal); err != nil {
				newErr := errors.New("The environment variable: GANDALF_" + keyName + " cannot be parsed as an integer using the value: " + strVal)
				fmt.Println(newErr)
				log.Fatalf("error while parsing an environment variable: %v",newErr)
			}
		}
		if len(strVal) > 0 && *(keyDef.value) == "" {
			*(keyDef.value) = strVal
		}
	}
}

func yamlFileToMap() map[interface{}]map[interface{}]string {
	//Set a map from config yaml file
	keyDef := ConfigKeys["config_file"]
	if *(keyDef.value) == "" {
		*(keyDef.value) = keyDef.defaultVal
	}

	yamlMap := make(map[interface{}]map[interface{}]string)
	yamlFile, err := ioutil.ReadFile(*(keyDef.value))
	err = yaml.Unmarshal(yamlFile, &yamlMap)
	if err != nil {
		log.Fatalf("error while parsing the file: %v", err)
	}
	return yamlMap
}

func yamlFileParse() {
	//Parse the yaml parameters
	tempMap := yamlFileToMap()
	for keyName := range ConfigKeys {
		keyDef := ConfigKeys[keyName]
		if *(keyDef.value) == "" {
			*(keyDef.value) = tempMap[keyDef.component][keyName]
		}
		if keyDef.valType == "integer" && *(keyDef.value) != "" {
			if _, err := strconv.Atoi(*(keyDef.value)); err != nil {
				newErr := errors.New("The Yaml parameter for: " + keyName + " cannot be parsed as an integer using the value: " + *(keyDef.value))
				log.Fatalf("error while parsing a Yaml parameter: %v",newErr)
			}
		}
	}
}

func defaultParse() {
	// parse default values
	for keyName := range ConfigKeys {
		keyDef := ConfigKeys[keyName]
		if *(keyDef.value) == "" {
			*(keyDef.value) = keyDef.defaultVal
		}
	}
}

func printCfKeys() error {
	for keyName := range ConfigKeys {
		keyDef := ConfigKeys[keyName]
		componentType := keyDef.component
		if keyDef.valType == "string" {
			strVal, err := GetStringConfig(keyName)
			if err != nil {
				return err
			}
			fmt.Println(componentType + ": " + keyName + ": " + strVal)
		} else {
			intVal, err := GetIntegerConfig(keyName)
			if err != nil {
				return err
			}
			fmt.Println(componentType + ": " + keyName + ": " + strconv.Itoa(intVal))
		}
	}
	return nil
}

func parseConfig() {
	argParse()
	tempEnvVarSet()
	envParse()
	yamlFileParse()
	defaultParse()
}

func ConfigMain() {
	InitMainConfigKeys()
	parseConfig()
	_ = printCfKeys()
}
