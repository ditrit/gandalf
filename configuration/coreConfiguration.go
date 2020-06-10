package configuration



var (
	CoreConfigKeys = ConfigKeys
)

func InitCoreKeys(){
	_ = SetStringKeyConfig("core", "config_file", "f", "/etc/gandalf/gandalf.yaml", "path to the configuration file")
	_ = SetStringKeyConfig("core","logical_name","l","logical name","logical name of the component")
	_ = SetStringKeyConfig("core", "gandalf_type","g","gandalf type","launch mode (connector|aggregator|cluster)")
	_ = SetStringKeyConfig("core","cert_pem","","/etc/gandalf/cert/cert.pem","path of the TLS certificate")
	_ = SetStringKeyConfig("core","key_pem","","/etc/gandalf/cert/key.pem","path of the TLS private key")
}
/*
CoreConfigKeys["core"] = map[string]ConfigKey{
"config_file":  {"string", "chemin vers le fichier le configuration", "/etc/gandalf/gandalf.yaml", nil, "f", false, nil},
"logical_name": {"string", "nom logique", "", nil, "l", true, nil},
"gandalf_type": {"string", "mode de démarrage (connector|aggregator|cluster)", "", []string{"connector", "aggregator", "cluster"}, "g", true, nil},
"cert_pem":     {"string", "chemin vers le certificat TLS", "/etc/gandalf/cert/cert.pem", nil, "", false, nil},
"key_pem":      {"string", "chemin vers la clef privée TLS", "/etc/gandalf/cert/key.pem", nil, "", false, nil},
"test": {"string", "test", "", nil, "", false, nil},
}*/


func yamlConfiguration(){

}

func defaultConfiguration(){

}




/*
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