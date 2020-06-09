package configuration


import (
"errors"
"flag"
"fmt"
"strconv"
)

type configKey struct {
	value   *string
	Type 	string
	Default string
	Usage   string
}


var ConfigResults = make(map[string]map[string]configKey)

func initConfig() {
	ConfigResults["core"] = map[string]configKey{
		"testFlag": { nil, "string", "testFlag", "testFlag "},
		"test2": {nil, "integer", "test2", "test2"},
	}
	ConfigResults["connectors"] = map[string]configKey{
		"connectorFlag": {nil, "string", "testFlag", "testFlag "},
		"Connector2": {nil, "integer", "test2", "test2"},
	}
}



func getStringConfig(keyName string) (string,error) {
	k,exists := ConfigResults["core"][keyName]
	if !exists {
		k,exists = ConfigResults["connectors"][keyName]
	}
	if exists {
		if k.Type == "string" {
			return *k.value, nil
		} else {
			return "", errors.New("This key "+ keyName +" is of type " + k.Type + ", not string")
		}
	}else{
		return "",errors.New("The key "+ keyName +" doesn't exist ")
	}
}

func getIntegerConfig(keyName string) (int64,error){
	k,exists := ConfigResults["core"][keyName]
	if !exists {
		k,exists = ConfigResults["connectors"][keyName]
	}
	if exists {
		if k.Type == "integer" {
			return strconv.ParseInt(*k.value,10,64)
		} else {
			return -1, errors.New("This key "+ keyName +" is of type " + k.Type + ", not integer")
		}
	}else{
		return -1,errors.New("The key "+ keyName +"doesn't exist ")
	}
}

func main() {
	initConfig()
	for gandalfType, _ := range ConfigResults {
		for key, _ := range ConfigResults[gandalfType] {
			flag.StringVar(ConfigResults[gandalfType][key].value, key, ConfigResults[gandalfType][key].Default, ConfigResults[gandalfType][key].Usage)
		}
	}
	flag.Parse()
	for gandalfType, _ := range ConfigResults {
		for key, _ := range ConfigResults[gandalfType] {
			fmt.Print(key,": ")
			if ConfigResults[gandalfType][key].Type == "string"{
				fmt.Println(getStringConfig(key))
			}else{
				fmt.Println(getIntegerConfig(key))
			}

		}
	}

}
