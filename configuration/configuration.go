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
		return -1, errors.New("The value '" + *keyDef.value + "' privided for the key '" + keyName + "' can not be parsed as an integer.")
	}
	return intValue, nil
}

func InitMainConfigKeys() {
	InitCfgKeys()
	InitCoreKeys()
}

func InitCfgKeys() {
	_ = SetStringKeyConfig("TEST", "testFlag", "flag", "", "a string test paramater for core")
	_ = SetIntegerKeyConfig("TEST", "test2", "", 9, "an integer test parameter for core")
	_ = SetStringKeyConfig("TEST", "connectorFlag", "v", "", "a string parameter for connectors")
	_ = SetIntegerKeyConfig("TEST", "connector2", "w", 22, "an integer parameter for connectors")
}

func InitCoreKeys() {
	_ = SetStringKeyConfig("core", "config_file", "f", "configuration/elements/gandalf.yaml", "path to the configuration file")
	_ = SetStringKeyConfig("core", "logical_name", "l", "logical name", "logical name of the component")
	_ = SetStringKeyConfig("core", "gandalf_type", "g", "gandalf type", "launch mode (connector|aggregator|cluster)")
	_ = SetStringKeyConfig("core", "cert_pem", "", "/etc/gandalf/cert/cert.pem", "path of the TLS certificate")
	_ = SetStringKeyConfig("core", "key_pem", "", "/etc/gandalf/cert/key.pem", "path of the TLS private key")
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
}

func tempEnvVarSet() {
	//temporary environment variables setter
	_ = os.Setenv("GANDALF_connectorFlag", "testENV")
	_ = os.Setenv("GANDALF_connector2", "12")
}

func envParse() error {
	// parse environment variables
	for keyName := range ConfigKeys {
		keyDef := ConfigKeys[keyName]
		strVal := os.Getenv("GANDALF_" + keyName)
		/*if keyDef.valType == "integer"{
			_, err := strconv.Atoi(strVal)
			println(err)
			if err != nil {
				newErr := errors.New("The environment variable: GANDALF_" + keyName + " cannot be parsed as an integer" )
				fmt.Println(newErr)
				return newErr
			}
		}*/
		if len(strVal) > 0 && *(keyDef.value) == "" {
			*(keyDef.value) = strVal
		}
	}
	return nil
}

func yamlFileToMap() map[interface{}]map[interface{}]string {
	keyDef := ConfigKeys["config_file"]
	filePath := ""
	if *(keyDef.value) == "" {
		filePath = keyDef.defaultVal
	} else {
		filePath = *(keyDef.value)
	}

	yamlMap := make(map[interface{}]map[interface{}]string)
	yamlFile, err := ioutil.ReadFile(filePath)
	err = yaml.Unmarshal(yamlFile, &yamlMap)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return yamlMap
}

func yamlFileParse() {
	tempMap := yamlFileToMap()
	for keyName := range ConfigKeys {
		keyDef := ConfigKeys[keyName]
		if *(keyDef.value) == "" {
			*(keyDef.value) = tempMap[keyDef.component][keyName]
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
	_ = envParse()
	yamlFileParse()
	defaultParse()
}

func ConfigMain() {
	InitMainConfigKeys()
	parseConfig()
	_ = printCfKeys()
}
