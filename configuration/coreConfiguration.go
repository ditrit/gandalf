package configuration

import (
	"flag"
	"fmt"
	"os"
)

type ConfigKey struct {
	Type          string
	Description   string
	Default       interface{}
	AllowedValues []string
	ShortName     string
	Mandatory     bool
	Value         *string
}

var (
	CoreConfigKeys = map[string]map[string]ConfigKey{}
)

func initConfig(){
	CoreConfigKeys["core"] = map[string]ConfigKey{
		"config_file":  {"string", "chemin vers le fichier le configuration", "/etc/gandalf/gandalf.yaml", nil, "f", false, nil},
		"logical_name": {"string", "nom logique", "", nil, "l", true, nil},
		"gandalf_type": {"string", "mode de démarrage (connector|aggregator|cluster)", "", []string{"connector", "aggregator", "cluster"}, "g", true, nil},
		"cert_pem":     {"string", "chemin vers le certificat TLS", "/etc/gandalf/cert/cert.pem", nil, "", false, nil},
		"key_pem":      {"string", "chemin vers la clef privée TLS", "/etc/gandalf/cert/key.pem", nil, "", false, nil},
		"test": {"string", "test", "", nil, "", false, nil},
	}
}


func SetConfiguration() {
	initConfig()
	envConfiguration()



	for gtype,_ := range CoreConfigKeys {
		for key,_ := range CoreConfigKeys[gtype] {
			fmt.Println(key,": ",CoreConfigKeys[gtype][key].Value)
		}
	}

}

func flagConfiguration(){
	// recuperer les valeurs passées en paramètre
	for gandalfType, _ := range CoreConfigKeys {
		for key, _ := range CoreConfigKeys[gandalfType] {
			flag.StringVar(CoreConfigKeys[gandalfType][key].Value,key, "test", CoreConfigKeys[gandalfType][key].Description)
			if CoreConfigKeys[gandalfType][key].ShortName != "" {
				flag.StringVar(CoreConfigKeys[gandalfType][key].Value, CoreConfigKeys[gandalfType][key].ShortName, "toto", CoreConfigKeys[gandalfType][key].Description)
			}
		}
	}
	flag.Parse()
}

func envConfiguration(){
	// recuperer les valeurs passées en variables d'envrionnement si pas de valeur déjà passée en paramètre
	os.Setenv("GANDALF_test","testEnv")
	for gandalfType, _ := range CoreConfigKeys {
		for key, _ := range CoreConfigKeys[gandalfType] {
			stringVal := os.Getenv("GANDALF_"+ key)
			fmt.Println(stringVal)
			if len(stringVal) > 0 {
					if CoreConfigKeys[gandalfType][key].Value == nil {
						*CoreConfigKeys[gandalfType][key].Value = stringVal
					}
				}
			}
		}
}

func yamlConfiguration(){

}

func defaultConfiguration(){

}




/*
// recuperer les valeurs passées en variables d'envrionnement si pas de valeur déjà passée en paramètre
for gtype, _ = range ConfigKeys {
for key, _ := range ConfigKeys[gtype] {
strval := os.Getenv( "GANDALF_" + key )
if len(strval) > 0 {
switch ConfigKeys[gtype][key].Type {
case "string":
if ConfigKeys[gtype][key].Value.(string) == "" {
ConfigKeys[gtype][key].Value.(string) = strval
}
case "int":
if ConfigKeys[gtype][key].Value.(int) == -1 {
ConfigKeys[gtype][key].Value.(int) = strconv.ParseInt((strval, 10, 64)
}
}
}
}
config_file := ConfigKeys["core"]["config_file"]
// recuperer les valeurs du fichier de configuration si pas de valeur déjà obtenue
nodes := yamlv3.Unmarshal(config_file)
for gtype, _ = range ConfigKeys {
for key, _ := range ConfigKeys[gtype] {
// si key est dans nodes et pas et si ConfigKeys[gtype][key].Value pas définie ("" pour string ou -1 pour int) alors affecter
}
}
// recuperer les valeurs par defaut si pas de valeur déjà obtenue et verifier si allowed value
for gtype, _ = range ConfigKeys {
for key, _ := range ConfigKeys[gtype] {
switch ConfigKeys[gtype][key].Type {
case "string":
if ConfigKeys[gtype][key].Value.(string) == "" && ConfigKeys[gtype][key].Default != nil { // et verifier si allowed value
ConfigKeys[gtype][key].Value.(string) = ConfigKeys[gtype][key].Default
}
case "int":
if ConfigKeys[gtype][key].Value.(int) == -1 && ConfigKeys[gtype][key].Default != nil { // et verifier si allowed value
ConfigKeys[gtype][key].Value.(int) = ConfigKeys[gtype][key].Default
}
}
}
}
var ConfigVals = make(map[string]interface{})
// valider que les valeurs obligatoires sont fournies
gandalf_type := configKeys["core"]["gandalf_type"].Value.(string)
if gandalf_type == "" {
Error( "grosse colère" )
}
for gtype, _ = range []string{ "core", gandalf_type } {
for key, _ := range ConfigKeys[gtype] {
if ConfigKeys[gtype][key].Mandatory == true {
switch ConfigKeys[gtype][key].Type {
case "string":
if ConfigKeys[gtype][key].Value.(string) == "" {
Error("grosse colère")
}
else {
ConfigVals[key] = ConfigKeys[gtype][key].Value
}
case "int":
if ConfigKeys[gtype][key].Value.(int) == -1 {
Error("grosse colère")
} else {
ConfigVals[key] = ConfigKeys[gtype][key].Value
}
}
}
}
}
*/