package api

// Urls :
type Urls struct {
	STATIC_PATH  string
	ROOT_PATH    string
	GANDALF_PATH string
	TENANTS_PATH string

	LOGIN_PATH                  string
	CLI_PATH                    string
	CLUSTER_PATH                string
	CLUSTER_PATH_LIST           string
	CLUSTER_PATH_CREATE         string
	CLUSTER_PATH_DECLARE_MEMBER string
	CLUSTER_PATH_READ           string
	CLUSTER_PATH_UPDATE         string
	CLUSTER_PATH_DELETE         string
	USER_PATH                   string
	USER_PATH_LIST              string
	USER_PATH_CREATE            string
	USER_PATH_READ_BY_NAME      string
	USER_PATH_READ              string
	USER_PATH_UPDATE            string
	USER_PATH_DELETE            string
	TENANT_PATH                 string
	TENANT_PATH_LIST            string
	TENANT_PATH_CREATE          string
	TENANT_PATH_READ            string
	TENANT_PATH_UPDATE          string
	TENANT_PATH_DELETE          string
	CONFIGURATION_PATH          string
	CONFIGURATION_PATH_LIST     string
	CONFIGURATION_PATH_CREATE   string
	CONFIGURATION_PATH_READ     string
	CONFIGURATION_PATH_UPDATE   string
	CONFIGURATION_PATH_DELETE   string
	CONFIGURATION_PATH_UPLOAD   string

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
	CONFIGURATION_TENANTS_PATH           string
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
	apiurls.GANDALF_PATH = "/gandalf"
	apiurls.TENANTS_PATH = "/tenants/{tenant}"

	//GANDALF
	apiurls.LOGIN_PATH = apiurls.GANDALF_PATH + "/login/"
	apiurls.CLI_PATH = apiurls.GANDALF_PATH + "/cli/"
	apiurls.CLUSTER_PATH = apiurls.GANDALF_PATH + "/clusters"
	apiurls.CLUSTER_PATH_LIST = apiurls.CLUSTER_PATH + "/"
	apiurls.CLUSTER_PATH_CREATE = apiurls.CLUSTER_PATH + "/"
	apiurls.CLUSTER_PATH_DECLARE_MEMBER = apiurls.CLUSTER_PATH + "/declare/"
	apiurls.CLUSTER_PATH_READ = apiurls.CLUSTER_PATH + "/{id}"
	apiurls.CLUSTER_PATH_UPDATE = apiurls.CLUSTER_PATH + "/{id:[0-9]+}"
	apiurls.CLUSTER_PATH_DELETE = apiurls.CLUSTER_PATH + "/{id:[0-9]+}"
	apiurls.USER_PATH = apiurls.GANDALF_PATH + "/users"
	apiurls.USER_PATH_LIST = apiurls.USER_PATH + "/"
	apiurls.USER_PATH_CREATE = apiurls.USER_PATH + "/"
	apiurls.USER_PATH_READ = apiurls.USER_PATH + "/{id:[0-9]+}"
	apiurls.USER_PATH_READ_BY_NAME = apiurls.USER_PATH + "/{name}"
	apiurls.USER_PATH_UPDATE = apiurls.USER_PATH + "/{id:[0-9]+}"
	apiurls.USER_PATH_DELETE = apiurls.USER_PATH + "/{id:[0-9]+}"
	apiurls.TENANT_PATH = apiurls.GANDALF_PATH + "/tenants"
	apiurls.TENANT_PATH_LIST = apiurls.TENANT_PATH + "/"
	apiurls.TENANT_PATH_CREATE = apiurls.TENANT_PATH + "/"
	apiurls.TENANT_PATH_READ = apiurls.TENANT_PATH + "/{id:[0-9]+}"
	apiurls.TENANT_PATH_UPDATE = apiurls.TENANT_PATH + "/{id:[0-9]+}"
	apiurls.TENANT_PATH_DELETE = apiurls.TENANT_PATH + "/{id:[0-9]+}"
	apiurls.CONFIGURATION_PATH = apiurls.GANDALF_PATH + "/configurations"
	apiurls.CONFIGURATION_PATH_LIST = apiurls.CONFIGURATION_PATH + "/"
	apiurls.CONFIGURATION_PATH_CREATE = apiurls.CONFIGURATION_PATH + "/"
	apiurls.CONFIGURATION_PATH_READ = apiurls.CONFIGURATION_PATH + "/{id:[0-9]+}"
	apiurls.CONFIGURATION_PATH_UPDATE = apiurls.CONFIGURATION_PATH + "/{id:[0-9]+}"
	apiurls.CONFIGURATION_PATH_DELETE = apiurls.CONFIGURATION_PATH + "/{id:[0-9]+}"
	apiurls.CONFIGURATION_PATH_UPLOAD = apiurls.CONFIGURATION_PATH + "/upload/"

	//TENANTS
	apiurls.CONNECTOR_PATH = apiurls.TENANTS_PATH + "/connectors"
	apiurls.CONNECTOR_PATH_LIST = apiurls.CONNECTOR_PATH + "/"
	apiurls.CONNECTOR_PATH_CREATE = apiurls.CONNECTOR_PATH + "/"
	apiurls.CONNECTOR_PATH_DECLARE_MEMBER = apiurls.CONNECTOR_PATH + "/declare/{name}"
	apiurls.CONNECTOR_PATH_READ = apiurls.CONNECTOR_PATH + "/{id:[0-9]+}"
	apiurls.CONNECTOR_PATH_UPDATE = apiurls.CONNECTOR_PATH + "/{id:[0-9]+}"
	apiurls.CONNECTOR_PATH_DELETE = apiurls.CONNECTOR_PATH + "/{id:[0-9]+}"
	apiurls.AGGREGATOR_PATH = apiurls.TENANTS_PATH + "/aggregators"
	apiurls.AGGREGATOR_PATH_LIST = apiurls.AGGREGATOR_PATH + "/"
	apiurls.AGGREGATOR_PATH_CREATE = apiurls.AGGREGATOR_PATH + "/"
	apiurls.AGGREGATOR_PATH_DECLARE_MEMBER = apiurls.AGGREGATOR_PATH + "/declare/{name}"
	apiurls.AGGREGATOR_PATH_READ = apiurls.AGGREGATOR_PATH + "/{id:[0-9]+}"
	apiurls.AGGREGATOR_PATH_UPDATE = apiurls.AGGREGATOR_PATH + "/{id:[0-9]+}"
	apiurls.AGGREGATOR_PATH_DELETE = apiurls.AGGREGATOR_PATH + "/{id:[0-9]+}"
	apiurls.ROLE_PATH = apiurls.TENANTS_PATH + "/roles"
	apiurls.ROLE_PATH_LIST = apiurls.ROLE_PATH + "/"
	apiurls.ROLE_PATH_CREATE = apiurls.ROLE_PATH + "/"
	apiurls.ROLE_PATH_READ = apiurls.ROLE_PATH + "/{id:[0-9]+}"
	apiurls.ROLE_PATH_UPDATE = apiurls.ROLE_PATH + "/{id:[0-9]+}"
	apiurls.ROLE_PATH_DELETE = apiurls.ROLE_PATH + "/{id:[0-9]+}"
	apiurls.CONFIGURATION_TENANTS_PATH = apiurls.TENANTS_PATH + "/configurations"
	apiurls.CONFIGURATION_AGGREGATOR_PATH = apiurls.CONFIGURATION_TENANTS_PATH + "/aggregator"
	apiurls.CONFIGURATION_CONNECTOR_PATH = apiurls.CONFIGURATION_TENANTS_PATH + "/connector"
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
