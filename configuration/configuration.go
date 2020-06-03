package configuration

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

type config struct {
	configurationFilePath string
	logicalName string
	gandalfType string
	certPem string
	keyPem string
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

type FlagConfig struct {

}
var (
	gandalfConfig config

	defaultYamlPath = "configuration/elements/gandalf.yaml"

	YamlPathFlag string
	LogicalNameFlag string
 	GandalfTypeFlag string
	CertPemFlag string
	KeyPemFlag string
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
	flag.StringVar(&YamlPathFlag,"configfile","","set the path for the configuration file")
	flag.StringVar(&YamlPathFlag,"f","","set the path for the configuration file")
	flag.StringVar(&LogicalNameFlag,"lname","","Set the logical name of a component")
	flag.StringVar(&LogicalNameFlag,"l","","Set the logical name of a component")
	flag.StringVar(&GandalfTypeFlag,"gandalf_type","","Set the type of a component")
	flag.StringVar(&GandalfTypeFlag,"g","","Set the type of a component")
	flag.StringVar(&CertPemFlag,"cert_pem","","Set the certificate value")
	flag.StringVar(&KeyPemFlag,"key_pem","","Set the key value")
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



func PrintConfig() {

	fmt.Println("Configuration file path: ",gandalfConfig.configurationFilePath)
	fmt.Println("Logical name: ", gandalfConfig.logicalName)
	fmt.Println("Gandalf type: ",gandalfConfig.gandalfType)
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

func ChangeConfig(){
	gandalfConfig.setConfig()
}



func (c *yamlConfig) yamlToStruct(path string){
	yamlPath, err := ioutil.ReadFile(path)
	if err !=nil{
		fmt.Println("Error while finding the file:", err)
	}
	err = yaml.Unmarshal(yamlPath,c)
	if err !=nil{
		fmt.Printf("Unmarshal: %v\n", err)
	}
}

func(c *yamlConfig) getYamlConfig() *yamlConfig{
	if YamlPathFlag != "" {
		c.yamlToStruct(YamlPathFlag)
	} else {
		c.yamlToStruct(defaultYamlPath)
	}

	return c
}

func (c *config) setConfig() {
	yamlConfiguration := new(yamlConfig)
	yamlConfiguration.getYamlConfig()
	defaultPort := ":9800"

	if YamlPathFlag != ""{
		c.configurationFilePath = YamlPathFlag
	} else {
		c.configurationFilePath = defaultYamlPath
	}

	if LogicalNameFlag != "" {
		c.logicalName = LogicalNameFlag
	}else {
		c.logicalName = yamlConfiguration.LogicalName
	}

	if GandalfTypeFlag != "" {
		c.gandalfType = GandalfTypeFlag
	}else{
		c.gandalfType = yamlConfiguration.GandalfType
	}

	if CertPemFlag != "" {
		c.certPem = CertPemFlag
	}else{
		c.certPem = yamlConfiguration.CertPem
	}

	if KeyPemFlag != ""{
		c.keyPem = KeyPemFlag
	}else{
		c.keyPem = yamlConfiguration.KeyPem
	}

	if TenantFlag != "" {
		c.tenant = TenantFlag
	}else{
		c.tenant = yamlConfiguration.Tenant
	}

	if  CategoryFlag != "" {
		c.category = CategoryFlag
	}else{
		c.category = yamlConfiguration.Category
	}

	if ProductFlag != "" {
		c.product = ProductFlag
	}else{
		c.product = yamlConfiguration.Product
	}

	if AggregatorFlag != ""{
		c.aggregators = strings.Split(AggregatorFlag,",")
		for i ,_ := range c.aggregators {
			temp := strings.Split(c.aggregators[i],":")
			if len(temp) == 1{
				temp[0] += defaultPort
				c.aggregators[i] = temp[0]
			}
		}

	}else{
		c.aggregators = strings.Split(yamlConfiguration.Aggregators,",")
		for i ,_ := range c.aggregators {
			temp := strings.Split(c.aggregators[i],":")
				if len(temp) == 1{
					temp[0] += defaultPort
					c.aggregators[i] = temp[0]
				}
		}

	}

	if GandalfSecretFlag != "" {
		c.gandalfSecret = GandalfSecretFlag
	}else{
		c.gandalfSecret = yamlConfiguration.GandalfSecret
	}

	if ProductUrlFlag != "" {
		c.productUrl = strings.Split(ProductUrlFlag,",")
	}else{
		c.productUrl = strings.Split(yamlConfiguration.ProductUrl,",")
	}

	if LogPathFlag != ""{
		c.logPath = LogPathFlag
	}else{
		c.logPath = yamlConfiguration.LogPath
	}

	if MaxTimeoutFlag != 0 {
		c.maxTimeout = MaxTimeoutFlag
	}else {
		c.maxTimeout = yamlConfiguration.MaxTimeout
	}
}
