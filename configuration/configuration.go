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
	mandatory  bool    // is the value mandatory
}

var ConfigKeys = make(map[string]configKey)

// SetStringKeyConfig :
func SetStringKeyConfig(componentType string, keyName string, shortName string, defaultValue string, usage string, mandatory bool) error {
	keyDef, exists := ConfigKeys[keyName]
	if exists {
		return errors.New("The key " + keyName + "is already defined ( for component " + keyDef.component + ")")
	}
	ConfigKeys[keyName] = configKey{new(string), componentType, shortName, "string", defaultValue, usage, mandatory}
	return nil
}

// SetIntegerKeyConfig :
func SetIntegerKeyConfig(componentType string, keyName string, shortName string, defaultValue int, usage string, mandatory bool) error {
	keyDef, exists := ConfigKeys[keyName]
	if exists {
		return errors.New("The key " + keyName + "is already defined ( for component " + keyDef.component + ")")
	}
	ConfigKeys[keyName] = configKey{new(string), componentType, shortName, "integer", strconv.Itoa(defaultValue), usage, mandatory}
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
	InitTenantKey()
	InitConnectorKeys()
	InitAggregatorKeys()
	InitClusterKeys()
}

func InitCoreKeys() {
	_ = SetStringKeyConfig("core", "config_file", "f", "configuration/elements/gandalf.yaml", "path to the configuration file", true)
	_ = SetStringKeyConfig("core", "logical_name", "l", "", "logical name of the component", true)
	_ = SetStringKeyConfig("core", "gandalf_type", "g", "", "launch mode (connector|aggregator|cluster)", true)
	_ = SetStringKeyConfig("core", "cert_pem", "", "/etc/gandalf/cert/cert.pem", "path of the TLS certificate", false)
	_ = SetStringKeyConfig("core", "key_pem", "", "/etc/gandalf/cert/key.pem", "path of the TLS private key", false)
	_ = SetStringKeyConfig("core", "gandalf_log", "", "/etc/gandalf/log", "path of the log file", false)
}

func InitTenantKey(){
	_ = SetStringKeyConfig("connector/aggregator","tenant","t","","tenant of the component",true)
}

func InitConnectorKeys() {
	_ = SetStringKeyConfig("connector", "connector_type", "y", "svn", "category of the connector", true)
	_ = SetStringKeyConfig("connector", "product_name", "p", "product1", "product of the connector", true)
	_ = SetStringKeyConfig("connector", "aggregators", "a", "address1:9800,address2:6400,address3", "aggregators addresses linked to the connector", true)
	_ = SetStringKeyConfig("connector", "gandalf_secret", "s", "/etc/gandalf/gandalfSecret", "path of the gandalf secret", true)
	_ = SetStringKeyConfig("connector", "product_url", "u", "url1,url2,url3", "product url list of the connector", false)
	_ = SetStringKeyConfig("connector","workers","w","/etc/gandalf/workers","path for the workers configuration",false)
	_ = SetIntegerKeyConfig("connector", "max_timeout", "m", 100, "maximum timeout of the connector", false)
}

func InitAggregatorKeys() {
	_ = SetStringKeyConfig("aggregator", "clusters", "c", "address1[:9800],address2[:6300],address3", "clusters addresses linked to the aggregator", true)
}

func InitClusterKeys() {
	_ = SetStringKeyConfig("cluster", "cluster_join", "j", "clusterAddress", "cluster command (join)", true)
	_ = SetStringKeyConfig("cluster", "gandalf_db", "d", "pathToTheDB", "path for the gandalf database", false)
}

func argParse() error {
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
				fmt.Println("error while parsing a CLI parameter: ")
				return errors.New("The CLI parameter for: " + keyName + " cannot be parsed as an integer using the value: " + *(keyDef.value))
			}
		}
	}
	return nil
}

func envParse() error {
	// parse environment variables
	for keyName := range ConfigKeys {
		keyDef := ConfigKeys[keyName]
		strVal := os.Getenv("GANDALF_" + keyName)
		if keyDef.valType == "integer" && strVal != "" {
			if _, err := strconv.Atoi(strVal); err != nil {
				fmt.Println("error while parsing an environment variable: ")
				return errors.New("The environment variable: GANDALF_" + keyName + " cannot be parsed as an integer using the value: " + strVal)
			}
		}
		if len(strVal) > 0 && *(keyDef.value) == "" {
			*(keyDef.value) = strVal
		}
	}
	return nil
}

func yamlFileToMap() map[interface{}]map[interface{}]string {
	//Set a map from config yaml file
	keyDef := ConfigKeys["config_file"]
	if *(keyDef.value) == "" {
		*(keyDef.value) = keyDef.defaultVal
	}
	if _, err := os.Stat(*(keyDef).value); os.IsNotExist(err) {
		log.Fatalf("%v", err)
	}

	yamlMap := make(map[interface{}]map[interface{}]string)
	yamlFile, err := ioutil.ReadFile(*(keyDef.value))
	err = yaml.Unmarshal(yamlFile, &yamlMap)
	if err != nil {
		log.Fatalf("error while parsing the file: %v", err)
	}
	return yamlMap
}

func yamlFileParse() error {
	//Parse the yaml parameters
	tempMap := yamlFileToMap()
	for keyName := range ConfigKeys {
		keyDef := ConfigKeys[keyName]
		if *(keyDef.value) == "" {
			*(keyDef.value) = tempMap[keyDef.component][keyName]
		}
		if keyDef.valType == "integer" && *(keyDef.value) != "" {
			if _, err := strconv.Atoi(*(keyDef.value)); err != nil {
				fmt.Println("error while parsing a Yaml parameter: ")
				return errors.New("The Yaml parameter for: " + keyName + " cannot be parsed as an integer using the value: " + *(keyDef.value))
			}
		}
	}
	return nil
}

func defaultParse() error {
	// parse default values
	for keyName := range ConfigKeys {
		keyDef := ConfigKeys[keyName]
		if *(keyDef.value) == "" {
			*(keyDef.value) = keyDef.defaultVal
		}
	}
	return nil
}

func isConfigValid() error{
	gandalfType := *(ConfigKeys["gandalf_type"].value)
	if gandalfType != "connector" && gandalfType != "cluster" && gandalfType != "aggregator" {
		return errors.New("gandalf type isn't valid")
	}
	for keyName := range ConfigKeys {
		keyDef := ConfigKeys[keyName]
		if gandalfType == "connector" {
			if keyDef.component == "core" || keyDef.component == "connector" || keyDef.component == "connector/aggregator"{
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

func printCfKeys() error {
	for keyName := range ConfigKeys {
		keyDef := ConfigKeys[keyName]
		componentType := keyDef.component
		gandalfType := *(ConfigKeys["gandalf_type"].value)
		if gandalfType == "connector" {
			if componentType == "core" || componentType == "connector" || keyDef.component == "connector/aggregator"{
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
		}
		if gandalfType == "aggregator" {
			if componentType == "core" || componentType == "aggregator" || keyDef.component == "connector/aggregator" {
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
		}
		if gandalfType == "cluster" {
			if componentType == "core" || componentType == "cluster" {
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
		}
	}
	return nil
}

func ParseConfig() error {
	err := argParse()
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
	if err != nil {
		return err
	}
	return nil
}

func ConfigMain() {
	InitMainConfigKeys()
	err := ParseConfig()
	if err != nil {
		log.Fatalf("%v",err)
	}
	err = isConfigValid()
	if err != nil {
		log.Fatalf("%v",err)
	}
	_ = printCfKeys()
}
