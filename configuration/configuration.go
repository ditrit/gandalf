package configuration

import (
	"errors"
	"flag"
	"fmt"
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

func InitCfgKeys() {
	_ = SetStringKeyConfig("TEST core", "testFlag", "flag", "", "a string test paramater for core")
	_ = SetIntegerKeyConfig("TEST core", "test2", "", 9, "an integer test parameter for core")

	_ = SetStringKeyConfig("TEST connector", "connectorFlag", "v", "", "a string parameter for connectors")
	_ = SetIntegerKeyConfig("TEST connector", "connector2", "w", 22, "an integer parameter for connectors")
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

func tempEnvVarSet(){
	//temporary environment variables setter
	os.Setenv("GANDALF_connectorFlag","testENV")
	os.Setenv("GANDALF_connector2","25")
}

func envParse(){
	// parse environment variables
	for keyName := range ConfigKeys {
		keyDef := ConfigKeys[keyName]
		strVal := os.Getenv("GANDALF_"+ keyName)
		if len(strVal) > 0 && *(keyDef.value) == ""{
			*(keyDef.value) = strVal
		}
	}
}

func yamlFileParse(){

}

func defaultParse(){
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

func ConfigMain() {
	argParse()
	tempEnvVarSet()
	envParse()
	defaultParse()
	_ = printCfKeys()
}