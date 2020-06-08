package configuration

import (
	//"core/connector"
	"flag"
	"fmt"
	"io/ioutil"
)

type config struct {
	configurationFilePath string

	tenant string
	category string
	product string
	aggregators []string
	gandalfSecret string
	productUrl []string
	logPath string
	maxTimeout int64
}


type yamlConfig struct {
	LogicalName string `yaml:"logical_name"`
	GandalfType string `yaml:"gandalf_type"`
	CertPem string `yaml:"cert_pem"`
	KeyPem string `yaml:"key_pem"`
	Tenant string `yaml:"tenant"`
	Category string `yaml:"category"`
	Product string `yaml:"product"`
	Aggregators string `yaml:"aggregators"`
	GandalfSecret string `yaml:"gandalf_secret"`
	ProductUrl string `yaml:"product_url"`
	LogPath string `yaml:"log_path"`
	MaxTimeout int64 `yaml:"max_timeout"`
}

var (
	gandalfConfig config

	defaultYamlPath = "configuration/elements/gandalf.yaml"


	TenantFlag string
	CategoryFlag string
	ProductFlag string
	AggregatorFlag string
	GandalfSecretFlag string
	ProductUrlFlag string
	LogPathFlag string
	MaxTimeoutFlag int64


)


//flags to retrieve the settings of a component
func SetConnectorFlags(){
	flag.StringVar(&TenantFlag,"tenant","","Set the tenant of a component")
	flag.StringVar(&TenantFlag,"t","","Set the tenant of a component")
	flag.StringVar(&CategoryFlag,"category","","Set the category of a connector")
	flag.StringVar(&CategoryFlag,"c","","Set the category of a connector")
	flag.StringVar(&ProductFlag,"product","","Set the product of a connector")
	flag.StringVar(&ProductFlag,"p","","Set the product of a connector")
	flag.StringVar(&AggregatorFlag, "aggregators", "", "Set aggregator addresses")
	flag.StringVar(&AggregatorFlag,"a","","Set aggregator addresses")
	flag.StringVar(&GandalfSecretFlag,"gandalf_secret","","Set the path for the file with the secret")
	flag.StringVar(&GandalfSecretFlag,"s","","Set the path for the file with the secret")
	flag.StringVar(&ProductUrlFlag,"product_url","","Set the product url list")
	flag.StringVar(&ProductUrlFlag,"u","","Set the product url list")
	flag.StringVar(&LogPathFlag,"log","","Set the log path")
	flag.Int64Var(&MaxTimeoutFlag,"max_timeout",0,"Set the max timeout")
	flag.Int64Var(&MaxTimeoutFlag,"m",0,"Set the max timeout")

	flag.Parse()
}



func PrintConnectorConfig() {
	fmt.Println("Connector config:")
	fmt.Println("Configuration file path: ",gandalfConfig.configurationFilePath)
	fmt.Println("Tenant:",gandalfConfig.tenant)
	fmt.Println("Category:",gandalfConfig.category)
	fmt.Println("Product:",gandalfConfig.product)
	content, err := ioutil.ReadFile(gandalfConfig.gandalfSecret)
	if err != nil {
		fmt.Println("Error while finding the file:",err)
	}
	fmt.Println("Aggregators:",gandalfConfig.aggregators)
	fmt.Println("Gandalf secret:",gandalfConfig.gandalfSecret,"| secret:", string(content),)
	fmt.Println("Product url: ",gandalfConfig.productUrl)
	fmt.Println("Log path:",gandalfConfig.logPath)
	fmt.Println("Max timeout:",gandalfConfig.maxTimeout)


}


