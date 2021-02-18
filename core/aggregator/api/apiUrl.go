package api

// Urls :
type Urls struct {
	STATIC_PATH string
	ROOT_PATH   string
	PATH        string

	LOGIN_PATH                           string
	CONNECTOR_PATH                       string
	CONNECTOR_PATH_LIST                  string
	CONNECTOR_PATH_CREATE                string
	CONNECTOR_PATH_DECLARE_MEMBER        string
	CONNECTOR_PATH_READ                  string
	CONNECTOR_PATH_UPDATE                string
	CONNECTOR_PATH_DELETE                string
	AGGREGATOR_PATH                      string
	AGGREGATOR_PATH_LIST                 string
	AGGREGATOR_PATH_CREATE               string
	AGGREGATOR_PATH_DECLARE_MEMBER       string
	AGGREGATOR_PATH_READ                 string
	AGGREGATOR_PATH_UPDATE               string
	AGGREGATOR_PATH_DELETE               string
	ROLE_PATH                            string
	ROLE_PATH_LIST                       string
	ROLE_PATH_CREATE                     string
	ROLE_PATH_READ                       string
	ROLE_PATH_UPDATE                     string
	ROLE_PATH_DELETE                     string
	USER_PATH                            string
	USER_PATH_LIST                       string
	USER_PATH_CREATE                     string
	USER_PATH_READ                       string
	USER_PATH_UPDATE                     string
	USER_PATH_DELETE                     string
	CONFIGURATION_PATH                   string
	CONFIGURATION_AGGREGATOR_PATH        string
	CONFIGURATION_CONNECTOR_PATH         string
	CONFIGURATION_AGGREGATOR_PATH_LIST   string
	CONFIGURATION_AGGREGATOR_PATH_CREATE string
	CONFIGURATION_AGGREGATOR_PATH_READ   string
	CONFIGURATION_AGGREGATOR_PATH_UPDATE string
	CONFIGURATION_AGGREGATOR_PATH_DELETE string
	CONFIGURATION_AGGREGATOR_PATH_UPLOAD string
	CONFIGURATION_CONNECTOR_PATH_LIST    string
	CONFIGURATION_CONNECTOR_PATH_CREATE  string
	CONFIGURATION_CONNECTOR_PATH_READ    string
	CONFIGURATION_CONNECTOR_PATH_UPDATE  string
	CONFIGURATION_CONNECTOR_PATH_DELETE  string
	CONFIGURATION_CONNECTOR_PATH_UPLOAD  string
}

// ReturnURLS :
func ReturnURLS() *Urls {

	//BASE
	apiurls := new(Urls)
	apiurls.ROOT_PATH = "/"
	apiurls.PATH = "/gandalf"

	//TENANTS
	apiurls.LOGIN_PATH = apiurls.PATH + "/login/"
	apiurls.CONNECTOR_PATH = apiurls.PATH + "/connectors"
	apiurls.CONNECTOR_PATH_LIST = apiurls.CONNECTOR_PATH + "/"
	apiurls.CONNECTOR_PATH_CREATE = apiurls.CONNECTOR_PATH + "/"
	apiurls.CONNECTOR_PATH_DECLARE_MEMBER = apiurls.CONNECTOR_PATH + "/declare/{name}"
	apiurls.CONNECTOR_PATH_READ = apiurls.CONNECTOR_PATH + "/{id:[0-9]+}"
	apiurls.CONNECTOR_PATH_UPDATE = apiurls.CONNECTOR_PATH + "/{id:[0-9]+}"
	apiurls.CONNECTOR_PATH_DELETE = apiurls.CONNECTOR_PATH + "/{id:[0-9]+}"
	apiurls.AGGREGATOR_PATH = apiurls.PATH + "/aggregators"
	apiurls.AGGREGATOR_PATH_LIST = apiurls.AGGREGATOR_PATH + "/"
	apiurls.AGGREGATOR_PATH_CREATE = apiurls.AGGREGATOR_PATH + "/"
	apiurls.AGGREGATOR_PATH_DECLARE_MEMBER = apiurls.AGGREGATOR_PATH + "/declare/{name}"
	apiurls.AGGREGATOR_PATH_READ = apiurls.AGGREGATOR_PATH + "/{id:[0-9]+}"
	apiurls.AGGREGATOR_PATH_UPDATE = apiurls.AGGREGATOR_PATH + "/{id:[0-9]+}"
	apiurls.AGGREGATOR_PATH_DELETE = apiurls.AGGREGATOR_PATH + "/{id:[0-9]+}"
	apiurls.ROLE_PATH = apiurls.PATH + "/roles"
	apiurls.ROLE_PATH_LIST = apiurls.ROLE_PATH + "/"
	apiurls.ROLE_PATH_CREATE = apiurls.ROLE_PATH + "/"
	apiurls.ROLE_PATH_READ = apiurls.ROLE_PATH + "/{id:[0-9]+}"
	apiurls.ROLE_PATH_UPDATE = apiurls.ROLE_PATH + "/{id:[0-9]+}"
	apiurls.ROLE_PATH_DELETE = apiurls.ROLE_PATH + "/{id:[0-9]+}"
	apiurls.USER_PATH = apiurls.PATH + "/users"
	apiurls.USER_PATH_LIST = apiurls.USER_PATH + "/"
	apiurls.USER_PATH_CREATE = apiurls.USER_PATH + "/"
	apiurls.USER_PATH_READ = apiurls.USER_PATH + "/{id:[0-9]+}"
	apiurls.USER_PATH_UPDATE = apiurls.USER_PATH + "/{id:[0-9]+}"
	apiurls.USER_PATH_DELETE = apiurls.USER_PATH + "/{id:[0-9]+}"
	apiurls.CONFIGURATION_PATH = apiurls.PATH + "/configurations"
	apiurls.CONFIGURATION_AGGREGATOR_PATH = apiurls.CONFIGURATION_PATH + "/aggregator"
	apiurls.CONFIGURATION_CONNECTOR_PATH = apiurls.CONFIGURATION_PATH + "/connector"
	apiurls.CONFIGURATION_AGGREGATOR_PATH_LIST = apiurls.CONFIGURATION_AGGREGATOR_PATH + "/"
	apiurls.CONFIGURATION_AGGREGATOR_PATH_CREATE = apiurls.CONFIGURATION_AGGREGATOR_PATH + "/"
	apiurls.CONFIGURATION_AGGREGATOR_PATH_READ = apiurls.CONFIGURATION_AGGREGATOR_PATH + "/{id:[0-9]+}"
	apiurls.CONFIGURATION_AGGREGATOR_PATH_UPDATE = apiurls.CONFIGURATION_AGGREGATOR_PATH + "/{id:[0-9]+}"
	apiurls.CONFIGURATION_AGGREGATOR_PATH_DELETE = apiurls.CONFIGURATION_AGGREGATOR_PATH + "/{id:[0-9]+}"
	apiurls.CONFIGURATION_AGGREGATOR_PATH_UPLOAD = apiurls.CONFIGURATION_AGGREGATOR_PATH + "/upload"
	apiurls.CONFIGURATION_CONNECTOR_PATH_LIST = apiurls.CONFIGURATION_CONNECTOR_PATH + "/"
	apiurls.CONFIGURATION_CONNECTOR_PATH_CREATE = apiurls.CONFIGURATION_CONNECTOR_PATH + "/"
	apiurls.CONFIGURATION_CONNECTOR_PATH_READ = apiurls.CONFIGURATION_CONNECTOR_PATH + "/{id:[0-9]+}"
	apiurls.CONFIGURATION_CONNECTOR_PATH_UPDATE = apiurls.CONFIGURATION_CONNECTOR_PATH + "/{id:[0-9]+}"
	apiurls.CONFIGURATION_CONNECTOR_PATH_DELETE = apiurls.CONFIGURATION_CONNECTOR_PATH + "/{id:[0-9]+}"
	apiurls.CONFIGURATION_CONNECTOR_PATH_UPLOAD = apiurls.CONFIGURATION_CONNECTOR_PATH + "/upload"

	return apiurls
}
