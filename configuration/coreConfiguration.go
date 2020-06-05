package configuration

import (
	"flag"
	"fmt"
)

type configKey struct {
	configType    string
	description   string
	configDefault interface{}
	shortName     string
	mandatory     bool
}

type coreConfig struct {
	logicalName string
	gandalfType string
	certPem string
	keyPem string
}

var (
	coreConfigKeys map[string]configKey
	coreConfiguration coreConfig

	YamlPathFlag string
	LogicalNameFlag string
	GandalfTypeFlag string
	CertPemFlag string
	KeyPemFlag string
)

func PrintCoreConfig(){
	fmt.Println("Core config: ")
	fmt.Println("Logical name: ", coreConfiguration.logicalName)
	fmt.Println("Gandalf type: ",coreConfiguration.gandalfType)
	fmt.Println("cert_pem:", coreConfiguration.certPem)
	fmt.Println("key_pem:", coreConfiguration.keyPem)

	fmt.Println()
}

func SetCoreFlags(){
	flag.StringVar(&YamlPathFlag,"configfile","","set the path for the configuration file")
	flag.StringVar(&YamlPathFlag,"f","","set the path for the configuration file")
	flag.StringVar(&LogicalNameFlag,"lname","","Set the logical name of a component")
	flag.StringVar(&LogicalNameFlag,"l","","test")
	flag.StringVar(&GandalfTypeFlag,"gandalf_type","","Set the type of a component")
	flag.StringVar(&GandalfTypeFlag,"g","","Set the type of a component")
	flag.StringVar(&CertPemFlag,"cert_pem","","Set the certificate value")
	flag.StringVar(&KeyPemFlag,"key_pem","","Set the key value")

	flag.Parse()
}

func SetCoreConfig(){
	coreConfiguration.setConfig()
}

func (c *coreConfig) setConfig(){
	coreConfigKeys = make(map[string]configKey)

	if LogicalNameFlag != "" {
		c.logicalName = LogicalNameFlag
	}else {
		c.logicalName = "pouet"
	}

	if GandalfTypeFlag != "" {
		c.gandalfType = GandalfTypeFlag
	}else{
		c.gandalfType = "connector"
	}

	if CertPemFlag != "" {
		c.certPem = CertPemFlag
	}else{
		c.certPem = "test path"
	}

	if KeyPemFlag != ""{
		c.keyPem = KeyPemFlag
	}else{
		c.keyPem = "test path2"
	}
}
