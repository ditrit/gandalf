package api

type Urls struct {
	STATIC_PATH  string
	ROOT_PATH    string
	GANDALF_PATH string
	TENANTS_PATH string

	GANDALF_CLUSTER_PATH        string
	GANDALF_CLUSTER_PATH_LIST   string
	GANDALF_CLUSTER_PATH_CREATE string
	GANDALF_CLUSTER_PATH_READ   string
	GANDALF_CLUSTER_PATH_UPDATE string
	GANDALF_CLUSTER_PATH_DELETE string
	GANDALF_ROLE_PATH           string
	GANDALF_ROLE_PATH_LIST      string
	GANDALF_ROLE_PATH_CREATE    string
	GANDALF_ROLE_PATH_READ      string
	GANDALF_ROLE_PATH_UPDATE    string
	GANDALF_ROLE_PATH_DELETE    string
	GANDALF_USER_PATH           string
	GANDALF_USER_PATH_LIST      string
	GANDALF_USER_PATH_CREATE    string
	GANDALF_USER_PATH_READ      string
	GANDALF_USER_PATH_UPDATE    string
	GANDALF_USER_PATH_DELETE    string
	GANDALF_TENANT_PATH         string
	GANDALF_TENANT_PATH_LIST    string
	GANDALF_TENANT_PATH_CREATE  string
	GANDALF_TENANT_PATH_READ    string
	GANDALF_TENANT_PATH_UPDATE  string
	GANDALF_TENANT_PATH_DELETE  string

	TENANTS_CONNECTOR_PATH         string
	TENANTS_CONNECTOR_PATH_LIST    string
	TENANTS_CONNECTOR_PATH_CREATE  string
	TENANTS_CONNECTOR_PATH_READ    string
	TENANTS_CONNECTOR_PATH_UPDATE  string
	TENANTS_CONNECTOR_PATH_DELETE  string
	TENANTS_AGGREGATOR_PATH        string
	TENANTS_AGGREGATOR_PATH_LIST   string
	TENANTS_AGGREGATOR_PATH_CREATE string
	TENANTS_AGGREGATOR_PATH_READ   string
	TENANTS_AGGREGATOR_PATH_UPDATE string
	TENANTS_AGGREGATOR_PATH_DELETE string
	TENANTS_ROLE_PATH              string
	TENANTS_ROLE_PATH_LIST         string
	TENANTS_ROLE_PATH_CREATE       string
	TENANTS_ROLE_PATH_READ         string
	TENANTS_ROLE_PATH_UPDATE       string
	TENANTS_ROLE_PATH_DELETE       string
	TENANTS_USER_PATH              string
	TENANTS_USER_PATH_LIST         string
	TENANTS_USER_PATH_CREATE       string
	TENANTS_USER_PATH_READ         string
	TENANTS_USER_PATH_UPDATE       string
	TENANTS_USER_PATH_DELETE       string
}

func ReturnURLS() *Urls {

	//BASE
	apiurls := new(Urls)
	apiurls.ROOT_PATH = "/"
	apiurls.GANDALF_PATH = "/gandalf"
	apiurls.TENANTS_PATH = "/{tenant}"

	//GANDALF
	apiurls.GANDALF_CLUSTER_PATH = apiurls.GANDALF_PATH + "/cluster"
	apiurls.GANDALF_CLUSTER_PATH_LIST = apiurls.GANDALF_CLUSTER_PATH + "/list"
	apiurls.GANDALF_CLUSTER_PATH_CREATE = apiurls.GANDALF_CLUSTER_PATH + "/create"
	apiurls.GANDALF_CLUSTER_PATH_READ = apiurls.GANDALF_CLUSTER_PATH + "/read/{id}"
	apiurls.GANDALF_CLUSTER_PATH_UPDATE = apiurls.GANDALF_CLUSTER_PATH + "/update/{id:[0-9]+}"
	apiurls.GANDALF_CLUSTER_PATH_DELETE = apiurls.GANDALF_CLUSTER_PATH + "/delete/{id:[0-9]+}"
	apiurls.GANDALF_ROLE_PATH = apiurls.GANDALF_PATH + "/role"
	apiurls.GANDALF_ROLE_PATH_LIST = apiurls.GANDALF_ROLE_PATH + "/list"
	apiurls.GANDALF_ROLE_PATH_CREATE = apiurls.GANDALF_ROLE_PATH + "/create"
	apiurls.GANDALF_ROLE_PATH_READ = apiurls.GANDALF_ROLE_PATH + "/read/{id:[0-9]+}"
	apiurls.GANDALF_ROLE_PATH_UPDATE = apiurls.GANDALF_ROLE_PATH + "/update/{id:[0-9]+}"
	apiurls.GANDALF_ROLE_PATH_DELETE = apiurls.GANDALF_ROLE_PATH + "/delete/{id:[0-9]+}"
	apiurls.GANDALF_USER_PATH = apiurls.GANDALF_PATH + "/user"
	apiurls.GANDALF_USER_PATH_LIST = apiurls.GANDALF_USER_PATH + "/list"
	apiurls.GANDALF_USER_PATH_CREATE = apiurls.GANDALF_USER_PATH + "/create"
	apiurls.GANDALF_USER_PATH_READ = apiurls.GANDALF_USER_PATH + "/read/{id:[0-9]+}"
	apiurls.GANDALF_USER_PATH_UPDATE = apiurls.GANDALF_USER_PATH + "/update/{id:[0-9]+}"
	apiurls.GANDALF_USER_PATH_DELETE = apiurls.GANDALF_USER_PATH + "/delete/{id:[0-9]+}"
	apiurls.GANDALF_TENANT_PATH = apiurls.GANDALF_PATH + "/tenant"
	apiurls.GANDALF_TENANT_PATH_LIST = apiurls.GANDALF_TENANT_PATH + "/list"
	apiurls.GANDALF_TENANT_PATH_CREATE = apiurls.GANDALF_TENANT_PATH + "/create"
	apiurls.GANDALF_TENANT_PATH_READ = apiurls.GANDALF_TENANT_PATH + "/read/{id:[0-9]+}"
	apiurls.GANDALF_TENANT_PATH_UPDATE = apiurls.GANDALF_TENANT_PATH + "/update/{id:[0-9]+}"
	apiurls.GANDALF_TENANT_PATH_DELETE = apiurls.GANDALF_TENANT_PATH + "/delete/{id:[0-9]+}"

	//TENANTS
	apiurls.TENANTS_CONNECTOR_PATH = apiurls.TENANTS_PATH + "/connector"
	apiurls.TENANTS_CONNECTOR_PATH_LIST = apiurls.TENANTS_CONNECTOR_PATH + "/list"
	apiurls.TENANTS_CONNECTOR_PATH_CREATE = apiurls.TENANTS_CONNECTOR_PATH + "/create"
	apiurls.TENANTS_CONNECTOR_PATH_READ = apiurls.TENANTS_CONNECTOR_PATH + "/read/{id:[0-9]+}"
	apiurls.TENANTS_CONNECTOR_PATH_UPDATE = apiurls.TENANTS_CONNECTOR_PATH + "/update/{id:[0-9]+}"
	apiurls.TENANTS_CONNECTOR_PATH_DELETE = apiurls.TENANTS_CONNECTOR_PATH + "/delete/{id:[0-9]+}"
	apiurls.TENANTS_AGGREGATOR_PATH = apiurls.TENANTS_PATH + "/aggregator"
	apiurls.TENANTS_AGGREGATOR_PATH_LIST = apiurls.TENANTS_AGGREGATOR_PATH + "/list"
	apiurls.TENANTS_AGGREGATOR_PATH_CREATE = apiurls.TENANTS_AGGREGATOR_PATH + "/create"
	apiurls.TENANTS_AGGREGATOR_PATH_READ = apiurls.TENANTS_AGGREGATOR_PATH + "/read/{id:[0-9]+}"
	apiurls.TENANTS_AGGREGATOR_PATH_UPDATE = apiurls.TENANTS_AGGREGATOR_PATH + "/update/{id:[0-9]+}"
	apiurls.TENANTS_AGGREGATOR_PATH_DELETE = apiurls.TENANTS_AGGREGATOR_PATH + "/delete/{id:[0-9]+}"
	apiurls.TENANTS_ROLE_PATH = apiurls.TENANTS_PATH + "/role"
	apiurls.TENANTS_ROLE_PATH_LIST = apiurls.TENANTS_ROLE_PATH + "/list"
	apiurls.TENANTS_ROLE_PATH_CREATE = apiurls.TENANTS_ROLE_PATH + "/create"
	apiurls.TENANTS_ROLE_PATH_READ = apiurls.TENANTS_ROLE_PATH + "/read/{id:[0-9]+}"
	apiurls.TENANTS_ROLE_PATH_UPDATE = apiurls.TENANTS_ROLE_PATH + "/update/{id:[0-9]+}"
	apiurls.TENANTS_ROLE_PATH_DELETE = apiurls.TENANTS_ROLE_PATH + "/delete/{id:[0-9]+}"
	apiurls.TENANTS_USER_PATH = apiurls.TENANTS_PATH + "/user"
	apiurls.TENANTS_USER_PATH_LIST = apiurls.TENANTS_USER_PATH + "/list"
	apiurls.TENANTS_USER_PATH_CREATE = apiurls.TENANTS_USER_PATH + "/create"
	apiurls.TENANTS_USER_PATH_READ = apiurls.TENANTS_USER_PATH + "/read/{id:[0-9]+}"
	apiurls.TENANTS_USER_PATH_UPDATE = apiurls.TENANTS_USER_PATH + "/update/{id:[0-9]+}"
	apiurls.TENANTS_USER_PATH_DELETE = apiurls.TENANTS_USER_PATH + "/delete/{id:[0-9]+}"

	return apiurls
}
